#!/bin/bash

for d in $(find . -type d | sort); do 
    echo "--- $d ---"
    go doc "$d" 
done
