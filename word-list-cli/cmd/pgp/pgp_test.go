package pgp

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

// test vectors from https://en.wikipedia.org/wiki/PGP_word_list
func TestPgpEncoder(t *testing.T) {
	// test vectors from https://en.wikipedia.org/wiki/PGP_word_list
	in, _ := hex.DecodeString("e58294f2e9a227486e8b061b31cc528fd7fa3f19")
	expect := strings.Split("topmost Istanbul Pluto vagabond treadmill Pacific brackish dictator goldfish Medusa afflict bravado chatter revolver Dupont midsummer stopwatch whimsical cowbell bottomless", " ")
	testEncoder(t, in, expect)

	// test vectors from https://github.com/quchen/pgp-wordlist/blob/master/test/Main.hs
	in, _ = hex.DecodeString("b268fce42cf5e504b4fbece3c8e429a5")
	expect = strings.Split("sawdust gravity wayside tradition Burbank visitor topmost alkali scenic Wichita tumor torpedo spaniel tradition breakup paperweight", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("bb9668c633114f7c515e7fe31bc06fc7")
	expect = strings.Split("shamrock monument frighten responsive chisel Babylon dropper informant drunken finicky lockup torpedo beeswax recipe gremlin retraction", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("2d568c1195474f6fe40546fec9f4419a")
	expect = strings.Split("button escapade offload Babylon preclude determine dropper hemisphere tonic almighty cubic yesteryear spearhead Virginia cranky newsletter", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("e975738c03c1497ffdced38ffa3dd8f2")
	expect = strings.Split("treadmill impartial hockey megaton acme recover deckhand integrate willow sardonic stapler midsummer wallet crucifix stormy vagabond", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("721907f2ca6fe4ee528f473abdf3645b")
	expect = strings.Split("highchair bottomless ahead vagabond spellbind hemisphere tonic universe Dupont midsummer dashboard corrosion skullcap vertigo flytrap exodus", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("b28d47f88c57b5092e26f12465b9aed7")
	expect = strings.Split("sawdust microscope dashboard warranty offload Eskimo scorecard applicant buzzard caretaker unwind Capricorn fracture proximate robust stethoscope", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("b3c67564a79e0abb8498d333f1e42a92")
	expect = strings.Split("scallion responsive indulge getaway repay onlooker allow publisher mural narrative stapler concurrent unwind tradition brickyard misnomer", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("cee71cbfe6594948067ad2fcb7744818")
	expect = strings.Split("spyglass truncated befriend rebellion tracker examine deckhand dictator afflict infancy standard Wilmington seabird hydraulic deadbolt borderline", " ")
	testEncoder(t, in, expect)
}

func testEncoder(t *testing.T, input []byte, expected []string) {
	var out []string
	e := NewEncoder()

	for _, b := range input {
		word := e.GetWord(b)
		out = append(out, word)
	}

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("Output not correct. Expected: %v; Got: %v", expected, out)
	}
}

// test vectors from https://en.wikipedia.org/wiki/PGP_word_list
func TestPgpDecoder(t *testing.T) {
	in := strings.Split("topmost Istanbul Pluto vagabond treadmill Pacific brackish dictator goldfish Medusa afflict bravado chatter revolver Dupont midsummer stopwatch whimsical cowbell bottomless", " ")
	expected, _ := hex.DecodeString("e58294f2e9a227486e8b061b31cc528fd7fa3f19")
	testDecoder(t, in, expected)
}

func testDecoder(t *testing.T, input []string, expected []byte) {
	var out []byte
	decoder := NewDecoder()

	for _, v := range input {
		b, err := decoder.GetByte(v)
		if err == nil {
			out = append(out, b)
		} else {
			t.Fatalf("Output not correct: %v", err)
		}
	}

	if reflect.DeepEqual(expected, out) == false {
		t.Log(hex.Dump(out))
		t.Fatalf("Output not correct")
	}
}
