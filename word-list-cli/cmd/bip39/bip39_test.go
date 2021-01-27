package bip39

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

// test vectors from https://github.com/trezor/python-mnemonic/blob/master/vectors.json
func TestBip39Encoder(t *testing.T) {
	// 16 bytes
	in, _ := hex.DecodeString("00000000000000000000000000000000")
	expect := strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	expect = strings.Split("legal winner thank year wave sausage worth useful legal winner thank yellow", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("80808080808080808080808080808080")
	expect = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage above", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("ffffffffffffffffffffffffffffffff")
	expect = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong", " ")
	testEncoder(t, in, expect)

	// 24 bytes
	in, _ = hex.DecodeString("000000000000000000000000000000000000000000000000")
	expect = strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	expect = strings.Split("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("808080808080808080808080808080808080808080808080")
	expect = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffff")
	expect = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when", " ")
	testEncoder(t, in, expect)

	// 32 bytes
	in, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	expect = strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	expect = strings.Split("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("8080808080808080808080808080808080808080808080808080808080808080")
	expect = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	expect = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote", " ")
	testEncoder(t, in, expect)

	// Others
	in, _ = hex.DecodeString("9e885d952ad362caeb4efe34a8e91bd2")
	expect = strings.Split("ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("6610b25967cdcca9d59875f5cb50b0ea75433311869e930b")
	expect = strings.Split("gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("68a79eaca2324873eacc50cb9c6eca8cc68ea5d936f98787c60c7ebc74e6ce7c")
	expect = strings.Split("hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("c0ba5a8e914111210f2bd131f3d5e08d")
	expect = strings.Split("scheme spot photo card baby mountain device kick cradle pact join borrow", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("6d9be1ee6ebd27a258115aad99b7317b9c8d28b6d76431c3")
	expect = strings.Split("horn tenant knee talent sponsor spell gate clip pulse soap slush warm silver nephew swap uncle crack brave", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("9f6a2878b2520799a44ef18bc7df394e7061a224d2c33cd015b157d746869863")
	expect = strings.Split("panda eyebrow bullet gorilla call smoke muffin taste mesh discover soft ostrich alcohol speed nation flash devote level hobby quick inner drive ghost inside", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("23db8160a31d3e0dca3688ed941adbf3")
	expect = strings.Split("cat swing flag economy stadium alone churn speed unique patch report train", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")
	expect = strings.Split("light rule cinnamon wrap drastic word pride squirrel upgrade then income fatal apart sustain crack supply proud access", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")
	expect = strings.Split("all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("f30f8c1da665478f49b001d94c5fc452")
	expect = strings.Split("vessel ladder alter error federal sibling chat ability sun glass valve picture", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("c10ec20dc3cd9f652c7fac2f1230f7a3c828389a14392f05")
	expect = strings.Split("scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump", " ")
	testEncoder(t, in, expect)

	in, _ = hex.DecodeString("f585c11aec520db57dd353c69554b21a89b20fb0650966fa0a9d6f74fd989d8f")
	expect = strings.Split("void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold", " ")
	testEncoder(t, in, expect)
}

func testEncoder(t *testing.T, input []byte, expected []string) {
	var out []string
	bip := NewEncoder()

	for _, v := range input {
		word, ok := bip.GetWord(v)
		if ok {
			out = append(out, word)
		}
	}
	word, ok := bip.GetLastWord()
	if ok {
		out = append(out, word)
	}

	if reflect.DeepEqual(expected, out) == false {
		t.Fatalf("Output not correct")
	}
}

// test vectors from https://github.com/trezor/python-mnemonic/blob/master/vectors.json
func TestBip39Decoder(t *testing.T) {
	// 16 bytes
	in := strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about", " ")
	expected, _ := hex.DecodeString("00000000000000000000000000000000")
	testDecoder(t, in, expected)

	in = strings.Split("legal winner thank year wave sausage worth useful legal winner thank yellow", " ")
	expected, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	testDecoder(t, in, expected)

	in = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage above", " ")
	expected, _ = hex.DecodeString("80808080808080808080808080808080")
	testDecoder(t, in, expected)

	in = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong", " ")
	expected, _ = hex.DecodeString("ffffffffffffffffffffffffffffffff")
	testDecoder(t, in, expected)

	// 24 bytes
	in = strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent", " ")
	expected, _ = hex.DecodeString("000000000000000000000000000000000000000000000000")
	testDecoder(t, in, expected)

	in = strings.Split("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will", " ")
	expected, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	testDecoder(t, in, expected)

	in = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always", " ")
	expected, _ = hex.DecodeString("808080808080808080808080808080808080808080808080")
	testDecoder(t, in, expected)

	in = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when", " ")
	expected, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffff")
	testDecoder(t, in, expected)

	// 32 bytes
	in = strings.Split("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art", " ")
	expected, _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	testDecoder(t, in, expected)

	in = strings.Split("legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title", " ")
	expected, _ = hex.DecodeString("7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f")
	testDecoder(t, in, expected)

	in = strings.Split("letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless", " ")
	expected, _ = hex.DecodeString("8080808080808080808080808080808080808080808080808080808080808080")
	testDecoder(t, in, expected)

	in = strings.Split("zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote", " ")
	expected, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	testDecoder(t, in, expected)

	// Other
	in = strings.Split("ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic", " ")
	expected, _ = hex.DecodeString("9e885d952ad362caeb4efe34a8e91bd2")
	testDecoder(t, in, expected)

	in = strings.Split("gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog", " ")
	expected, _ = hex.DecodeString("6610b25967cdcca9d59875f5cb50b0ea75433311869e930b")
	testDecoder(t, in, expected)

	in = strings.Split("hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length", " ")
	expected, _ = hex.DecodeString("68a79eaca2324873eacc50cb9c6eca8cc68ea5d936f98787c60c7ebc74e6ce7c")
	testDecoder(t, in, expected)

	in = strings.Split("scheme spot photo card baby mountain device kick cradle pact join borrow", " ")
	expected, _ = hex.DecodeString("c0ba5a8e914111210f2bd131f3d5e08d")
	testDecoder(t, in, expected)

	in = strings.Split("horn tenant knee talent sponsor spell gate clip pulse soap slush warm silver nephew swap uncle crack brave", " ")
	expected, _ = hex.DecodeString("6d9be1ee6ebd27a258115aad99b7317b9c8d28b6d76431c3")
	testDecoder(t, in, expected)

	in = strings.Split("panda eyebrow bullet gorilla call smoke muffin taste mesh discover soft ostrich alcohol speed nation flash devote level hobby quick inner drive ghost inside", " ")
	expected, _ = hex.DecodeString("9f6a2878b2520799a44ef18bc7df394e7061a224d2c33cd015b157d746869863")
	testDecoder(t, in, expected)

	in = strings.Split("cat swing flag economy stadium alone churn speed unique patch report train", " ")
	expected, _ = hex.DecodeString("23db8160a31d3e0dca3688ed941adbf3")
	testDecoder(t, in, expected)

	in = strings.Split("light rule cinnamon wrap drastic word pride squirrel upgrade then income fatal apart sustain crack supply proud access", " ")
	expected, _ = hex.DecodeString("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")
	testDecoder(t, in, expected)

	in = strings.Split("all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform", " ")
	expected, _ = hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")
	testDecoder(t, in, expected)

	in = strings.Split("vessel ladder alter error federal sibling chat ability sun glass valve picture", " ")
	expected, _ = hex.DecodeString("f30f8c1da665478f49b001d94c5fc452")
	testDecoder(t, in, expected)

	in = strings.Split("scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump", " ")
	expected, _ = hex.DecodeString("c10ec20dc3cd9f652c7fac2f1230f7a3c828389a14392f05")
	testDecoder(t, in, expected)

	in = strings.Split("void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold", " ")
	expected, _ = hex.DecodeString("f585c11aec520db57dd353c69554b21a89b20fb0650966fa0a9d6f74fd989d8f")
	testDecoder(t, in, expected)
}

func testDecoder(t *testing.T, input []string, expected []byte) {
	var out []byte
	decoder := NewDecoder()

	for _, v := range input {
		b, err := decoder.GetBytes(v)
		if err == nil {
			out = append(out[:], b[:]...)
		}
	}

	_, err := decoder.CheckHash()
	if err != nil {
		t.Fatalf("Hash problem '%v'", err)
	}

	if reflect.DeepEqual(expected, out) == false {
		t.Log(hex.Dump(out))
		t.Fatalf("Output not correct")
	}
}
