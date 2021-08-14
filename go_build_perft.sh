#!/bin/bash

go build -o bin/sachista-chess-perft perft/main.go

env GOOS=windows GOARCH=amd64 go build -o bin/sachista-chess-perft.exe perft/main.go