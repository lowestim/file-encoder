#! /bin/bash

go build -o ./out/file-encoder ./src/main.go
cp ./out/file-encoder $HOME/bin/
