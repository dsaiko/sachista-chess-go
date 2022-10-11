#!/bin/bash

set -e

echo "This utility allows verifying perft results stored in perft.txt against stockfish results."
while IFS=\| read -r fen depth count
do

result=$(stockfish << EOF | grep "Nodes searched" | awk -F ':' '{print $2}'
position fen ${fen}
go perft ${depth}
EOF
)

if [ "${result}" -ne "${count}" ]; then
    echo "${fen}|${depth}|${count}|stockfish:${result}"
fi

done < perft.txt

