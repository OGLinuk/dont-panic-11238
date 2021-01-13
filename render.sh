#!/bin/bash

walk() {
    cd $*
    for thing in *; do
        if [[ -d $thing ]]; then
            walk $thing
        elif [[ $thing == "README.md" ]]; then
            pandoc $thing -o index.html
        else
            pandoc $thing -o ${thing%.*}.html
        fi
    done
    cd ..
}

# Change docs to $* if necessary
walk docs
