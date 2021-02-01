# word-list-cli Docker Image

## Description

Small docker image with `word-list-cli` tool.

`word-list-cli` is utility for encoding and decoding binary data to different mnemonic word lists. [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) and [PGP](https://en.wikipedia.org/wiki/PGP_word_list) algorithms are supported.

Mnemonic code or mnemonic sentences are a group of easy to remember words. A mnemonic code or sentence is superior for human interaction compared to the handling of raw binary or hexadecimal representations of data. The sentence could be written on paper or spoken over the telephone.

Although this utility supports encoding data of any length, it is rather practical for storing a small amount of data. For example to encode passwords, passphrases or cryptographic keys and print then on paper.

`word-list-cli` can be found at https://github.com/dhlavaty/word-list

## Usage

```sh
# for word-list-cli tool help
$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli
```

### BIP39 algorithm

simple string to BIP39 encoding

```sh
$ echo -n '123456789' | docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli bip39
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----
```

BIP39 encoding with input from file

```sh
$ cat input-file.bin
123456789

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli bip39 < input-file.bin
-----BEGIN BIP39-----
couple muscle snack heavy gloom orchard tone
-----END BIP39-----
```

BIP39 encoding with output to file

```sh
$ cat input-file.bin
123456789

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli bip39 < input-file.bin > output.txt

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

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli decodebip39 < output.txt > decoded-file.bin

$ cat decoded-file.bin
123456789

# double-check that binary files are same
$ shasum *.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  input-file.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  decoded-file.bin
```

### PGP algorithm

Encoding simple string to PGP word list

```sh
$ echo -n '123456789' | docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli pgp
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----
```

with input from file

```sh
$ cat input-file.bin
123456789

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli pgp < input-file.bin
-----BEGIN PGP WORD LIST-----
chatter component chisel confidence chopper congregate clamshell consulting classroom
-----END PGP WORD LIST-----
```

with output to file

```sh
$ cat input-file.bin
123456789

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli pgp < input-file.bin > output.txt

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

$ docker run --rm -i dhlavaty/word-list-cli:latest word-list-cli decodepgp < output.txt > decoded-file.bin

$ cat decoded-file.bin
123456789

# double-check that binary files are same
$ shasum *.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  input-file.bin
f7c3bc1d808e04732adf679965ccc34ca7ae3441  decoded-file.bin
```

## Docker build

Build image:

```sh
$ docker build . --tag dhlavaty/word-list-cli:latest
```

Lint dockerfile:

```sh
$ docker run --rm -i hadolint/hadolint < Dockerfile
```

## License

This project is licensed under MIT - http://opensource.org/licenses/MIT
