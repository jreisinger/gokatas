#!/bin/bash

DIR=$1
[[ "$DIR" == "" ]] && DIR="."

for d in $(find "$DIR" -maxdepth 1 -type d | sort); do 
    echo "--- $d ---"
    go doc "./$d" 
done
