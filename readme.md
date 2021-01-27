# word-list-cli

Utility for encoding and decoding binary data to different mnemonic world lists. _Currently only BIP39 is supported._

Mnemonic code or mnemonic sentences are a group of easy to remember words. A mnemonic code or sentence is superior for human interaction compared to the handling of raw binary or hexadecimal representations of data. The sentence could be written on paper or spoken over the telephone.

Although this utility supports encoding data of any length, it is rather practical for storing a small amount of data. For example to encode passwords, passphrases or cryptographic keys and print then on paper.

## BIP39

Encoding binary data to mnemonic BIP39 world list is **100% compatible** with BIP39 deterministic keys. For more info see https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

### Example

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

## BIP39+ specs

Unlike original BIP39 algorithm, used for deterministic crypto currency keys, `word-list-cli` tool supports encoding binary data of any length. We call it **BIP39+** (_BIP39 Plus_)

TODO more info

### Example of BIP39+

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

## Build

We build everything in Docker

```sh
# interactive shell with golang
$ docker run --rm -it -v "$PWD/word-list-cli":/usr/src/word-list-cli -w /usr/src/word-list-cli golang:latest

# build and run app
$ docker run --rm -t -v "$PWD/word-list-cli":/usr/src/word-list-cli -w /usr/src/word-list-cli golang:latest go run word-list-cli

$ echo -n '123' | go run word-list-cli bip39

$ echo -n 'couple muscles trumpet+abuse' | go run word-list-cli decodebip39 > out.bin

$ echo -n "123456789" | od -A n -t x1

# releasing

$ go mod tidy
$ go test ./...

# build final .tar.gz / .zip files
$ ./release.sh 0.0.1
```
