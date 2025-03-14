#!/bin/bash

build() {
	go build -o ./build/i18n
}

run() {
	go run . $1 $2
}

default() {
	run
}

add() {
	cat <<EOF > "pkg/i18n/i18n${1}.go"
package i18n

import (
	"fmt"
	"os"
)

func (p *Puzzles) Puzzle${1}(file *os.File) string {
	fmt.Println("Puzzle")
	return ""
}
EOF
	
	echo "Created puzzle file."
	echo "Adding test-input and input files"
	
	curl -s "https://i18n-puzzles.com/puzzle/${1}/test-input" -o "tests/test-input${1}.txt"

	curl -s "https://i18n-puzzles.com/puzzle/${1}/input" -o "tests/input${1}.txt"
}
"${@-:default}"
