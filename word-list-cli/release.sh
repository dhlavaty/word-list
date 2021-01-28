#!/bin/bash

# exit when any command fails
set -e

go test ./...

# default (in Docker) is 'linux amd64'
go clean -i
go build -ldflags "-X 'word-list-cli/cmd.Version=$1'" word-list-cli
tar -czvf word-list-cli-$1-linux-amd64.tar.gz ./word-list-cli

# for MacOS Intel
go clean -i
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'word-list-cli/cmd.Version=$1'" word-list-cli
tar -czvf word-list-cli-$1-macos-amd64.tar.gz ./word-list-cli

# for ARM (for example Raspberry Pi)
go clean -i
env GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-X 'word-list-cli/cmd.Version=$1'" word-list-cli
tar -czvf word-list-cli-$1-linux-arm64-raspberrypi.tar.gz ./word-list-cli

# for Windows amd64
apt update
apt-get install zip
go clean -i
env GOOS=windows GOARCH=amd64 go build -ldflags "-X 'word-list-cli/cmd.Version=$1'" word-list-cli
zip word-list-cli-$1-windows-amd64.zip ./word-list-cli.exe

go clean -i
echo "Finished OK"