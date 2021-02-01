# word-list-cli

Utility for encoding and decoding binary data to different mnemonic word lists. [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) and [PGP](https://en.wikipedia.org/wiki/PGP_word_list) algorithms are supported.

Mnemonic code or mnemonic sentences are a group of easy to remember words. A mnemonic code or sentence is superior for human interaction compared to the handling of raw binary or hexadecimal representations of data. The sentence could be written on paper or spoken over the telephone.

Although this utility supports encoding data of any length, it is rather practical for storing a small amount of data. For example to encode passwords, passphrases or cryptographic keys and print then on paper.

## BIP39 WORD LIST

Encoding binary data to mnemonic BIP39 word list is **100% compatible** with BIP39 deterministic keys. For more info see https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

#### Example

Encoding simple string to BIP39 word list

```sh
$ echo -n '123456789' | word-list-cli bip39
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----
```

with input from file

```sh
$ cat input-file.bin
123456789

$ word-list-cli bip39 < input-file.bin
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----
```

with output to file

```sh
$ cat input-file.bin
123456789

$ word-list-cli bip39 < input-file.bin > output.txt

$ cat output.txt
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----
```

Decoding BIP39 word list back to binary

```sh
$ cat output.txt
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----

$ word-list-cli decodebip39 < output.txt > decoded-file.bin

$ cat decoded-file.bin
123456789

# double-check that binary files are same
$ shasum *.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  input-file.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  decoded-file.bin
```

### BIP39+ specs

Unlike original BIP39 algorithm, used for deterministic crypto currency keys, `word-list-cli` tool supports encoding binary data of any length. We call it **BIP39+** (_BIP39 Plus_)

BIP39 was designed to support binary data's length that is a multiple of 32 bits. To support data of any length, we
came up with an extension, that is 100% backwards compatible with BIP39.

BIP39 encodes into 11 bits words, where input data is padded by checksum. If data length is a multiple of 32 bits, the checksum bits are not longer that 8 bits. In cases, when the checksum is 9, 10 or 11 bits long, we will add a `PLUS WORD` to our very last word. This will signalize, how many bits is in the checksum. For 9 bit checksum, it is a word "`+abuse`", as that is 9th word in BIP39 dictionary.

#### Example of BIP39+

```sh
$ echo -n '123' | word-list-cli bip39
-----BEGIN BIP39-----
couple muscle trumpet+abuse
-----END BIP39-----

$ echo -n '1234567' | word-list-cli bip39
-----BEGIN BIP39-----
couple muscle snack heavy gloom shift+access
-----END BIP39-----

$ echo -n '12345678901' | word-list-cli bip39
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tooth alert census+accident
-----END BIP39-----
```

## PGP WORD LIST

The PGP Word List ("[Pretty Good Privacy](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) word list") is simpler encoding. Unlike BIP39, it does not contain any checksum of original data. For more info see https://en.wikipedia.org/wiki/PGP_word_list

#### Example

Encoding simple string to PGP word list

```sh
$ echo -n '123456789' | word-list-cli pgp
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----
```

with input from file

```sh
$ cat input-file.bin
123456789

$ word-list-cli pgp < input-file.bin
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----
```

with output to file

```sh
$ cat input-file.bin
123456789

$ word-list-cli pgp < input-file.bin > output.txt

$ cat output.txt
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----
```

Decoding PGP word list back to binary

```sh
$ cat output.txt
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----

$ word-list-cli decodepgp < output.txt > decoded-file.bin

$ cat decoded-file.bin
123456789

# double-check that binary files are same
$ shasum *.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  input-file.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  decoded-file.bin
```

## Build

We build everything in Docker. Cheatsheet:

```sh
# interactive shell with golang
$ docker run --rm -it -v "$PWD/word-list-cli":/usr/src/word-list-cli -w /usr/src/word-list-cli golang:latest
# $ go test ./...
# $ gofmt -w .

# build and run app
$ docker run --rm -t -v "$PWD/word-list-cli":/usr/src/word-list-cli -w /usr/src/word-list-cli golang:latest go run word-list-cli

# BIP39
$ echo -n '123' | go run word-list-cli bip39

$ echo -n 'couple muscles trumpet+abuse' | go run word-list-cli decodebip39 > out.bin

# PGP
$ echo -n '123' | go run word-list-cli pgp

$ echo -n 'chatter component chisel' | go run word-list-cli decodepgp > out.bin

# TO HEX
$ echo -n "123456789" | od -A n -t x1

# releasing

$ go mod tidy
$ go test ./...

# build final .tar.gz / .zip files
$ ./release.sh 0.0.1
```
