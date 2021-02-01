package pgp

import (
	"fmt"
)

var (
	// wordList is the set of words to use.
	wordList []string

	// wordMap is a reverse lookup map for wordList.
	wordMap map[string]int
)

func init() {
	setWordList(WordList)
}

func setWordList(list []string) {
	wordList = list
	wordMap = map[string]int{}

	for i, v := range wordList {
		wordMap[v] = i
	}
}

// encoder type
type pgpe struct {
	byteCounter int
}

// returns new PGP encoder
func NewEncoder() pgpe {
	return pgpe{}
}

// Add one byte to word-list encoder and get encoded word.
func (p *pgpe) GetWord(b byte) (word string) {
	isOdd := p.byteCounter % 2
	p.byteCounter += 1

	return wordList[(int(b)*2)+isOdd]
}

// decoder type
type pgpd struct {
	byteCounter int
}

// returns new PGP decoder
func NewDecoder() pgpd {
	return pgpd{}
}

func (p *pgpd) GetByte(s string) (byte, error) {
	index, found := wordMap[s]
	if !found {
		return 0, fmt.Errorf("word `%v` not found in reverse map", s)
	}

	isOdd := p.byteCounter % 2
	p.byteCounter += 1

	isOdd2 := index % 2

	if isOdd != isOdd2 {
		return byte((index >> 1)), fmt.Errorf("even/odd check for word `%v` not correct", s)
	}

	return byte((index >> 1)), nil
}
