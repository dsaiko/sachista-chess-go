#!/bin/bash

go test saiko.cz/sachista/generator -bench=. -run=NOTHING -benchmem -memprofile=mem.prof -cpuprofile=cpu.prof -benchtime=20s

