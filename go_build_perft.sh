#!/bin/bash

go build -o bin/sachista-chess-perft perft/main.go
objdump -D bin/sachista-chess-perft > bin/sachista-chess-perft.objdump

env GOOS=windows GOARCH=amd64 go build -o bin/sachista-chess-perft.exe perft/main.go