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

"${@-:default}"
