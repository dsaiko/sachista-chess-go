#!/bin/bash

CGO_ENABLED=0 go build  -o bin/sachista-chess-perft perft/main.go
