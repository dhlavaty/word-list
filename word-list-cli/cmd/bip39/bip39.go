package bip39

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"math"
	"strings"
)

var (
	// wordList is the set of words to use.
	wordList []string

	// wordMap is a reverse lookup map for wordList.
	wordMap map[string]int
)

type bip39e struct {
	word        int
	bitCounter  int
	byteCounter int
	sum         hash.Hash
}

// returns new Bip39 encoder
func NewEncoder() bip39e {
	return bip39e{sum: sha256.New()}
}

// Add one byte to word-list encoder and get encoded word.
func (b *bip39e) GetWord(p byte) (word string, ok bool) {
	b.word = b.word << 8     // shift word
	b.word = b.word | int(p) // add new lower 8 bits
	b.bitCounter += 8        // add to counter, how may there are left
	b.byteCounter += 1

	b.sum.Write([]byte{p})

	if b.bitCounter < 11 {
		// if there is less then 11 bits, there is nothing to return
		return "", false
	}

	w := b.word >> (b.bitCounter - 11)
	w = w & 0b11111111111 // "zero-out" everything above lower 11 bits
	b.bitCounter -= 11

	return wordList[w], true
}

// Finish word-list generation and get last encoded word.
// This function MUST be called to properly finish encoding.
func (b *bip39e) GetLastWord() (word string, ok bool) {
	if b.bitCounter < 1 {
		// if there are no bits in buffer, we need to add 11 bits of hash
		b.bitCounter = 0
	}

	// we need (11 - bitCounter) bits from SHA256 as padding
	hashBitsNeeded := (11 - b.bitCounter)

	// we make place for hash
	b.word = (b.word << hashBitsNeeded) & 0b11111111111 // "zero-out" everything above lower 11 bits

	// prepare first 16bits from hash
	s := b.sum.Sum(nil)
	hashWord := (int(s[0]) << 8) | int(s[1])

	// now we keep only those bits, that we really need
	hashWord = hashWord >> (16 - hashBitsNeeded)

	// and finally add hash bits to buffer
	w := b.word | hashWord

	if hashBitsNeeded > 8 {
		// when hash is more than 8 bits, we add one PLUS word more
		return wordList[w] + "+" + wordList[hashBitsNeeded], true
	}

	return wordList[w], true
}

type bip39d struct {
	word       int
	bitCounter int
	sum        hash.Hash
}

func NewDecoder() bip39d {
	return bip39d{sum: sha256.New()}
}

func (b *bip39d) GetBytes(i string) ([]byte, error) {
	split := strings.Split(i, "+")
	i = split[0]

	hashBits := 4092 // to prevent nil and simplify algorithm; 4092 does not exist in wordlist
	if len(split) > 1 {
		index, found := wordMap[split[1]]
		if !found {
			return nil, fmt.Errorf("hash word `%v` not found in reverse map", i)
		}
		hashBits = index
	}

	index, found := wordMap[i]
	if !found {
		return nil, fmt.Errorf("word `%v` not found in reverse map", i)
	}

	b.word = ((b.word << 11) | index)
	b.bitCounter += 11

	var out []byte

	for b.bitCounter > 8 {
		if b.bitCounter == hashBits {
			// all bits that are now in our buffer belongs to HASH
			break
		}
		out = append(out, byte((b.word >> (b.bitCounter - 8))))
		b.bitCounter -= 8
	}

	b.sum.Write(out)
	return out, nil
}

// Finish word-list decoding - check hash.
// Returns full SHA256 hash of returned bytes,
// and checks if it matches with encoded one.
// This function MAY be called at the end of decoding process.
func (b *bip39d) CheckHash() ([]byte, error) {
	hash := b.sum.Sum(nil)

	if b.bitCounter <= 0 {
		return hash, nil // nil, fmt.Errorf("no hash bits found")
	}

	// prepare first 16bits from hash
	hashWord := (int(hash[0]) << 8) | int(hash[1])
	// now we keep only those bits, that we really need
	hashWord = hashWord >> (16 - b.bitCounter)

	// zero-out all bits except hash (aka keep last b.bitCounter's bits)
	b.word = b.word & int(powInt(2, b.bitCounter)-1)

	if b.word == hashWord {
		return hash, nil
	}

	return hash, fmt.Errorf("hash does not match (bip39) 0b%b !== 0b%b (sha256)", b.word, hashWord)
}

func init() {
	setWordList(English)
}

func setWordList(list []string) {
	wordList = list
	wordMap = map[string]int{}

	for i, v := range wordList {
		wordMap[v] = i
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
