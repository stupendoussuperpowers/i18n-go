package main

import (
	"fmt"
	"github.com/stupendoussuperpowers/i18n-go.git/pkg/i18n"
	"os"
	"strconv"
)

type puzzleFunction func(*os.File) string

func main() {
	funcs := []puzzleFunction{i18n.PuzzleOne, i18n.PuzzleTwo, i18n.PuzzleThree, i18n.PuzzleFour}

	puzzNumber, _ := strconv.Atoi(os.Args[1])

	if puzzNumber <= 0 {
		fmt.Println("Puzzle number ", puzzNumber, " requires going back in time :(")
		os.Exit(1)

	}

	if puzzNumber > len(funcs) {
		fmt.Println("Puzzle number ", puzzNumber, " isn't yet solved!")
		os.Exit(1)
	}

	file, err := os.Open(fmt.Sprintf("tests/input%d.txt", puzzNumber))
	if err != nil {
		fmt.Printf("Input file not found. @ tests/input%d.txt", puzzNumber)
		os.Exit(1)
	}

	defer file.Close()

	// os.Stdin = file

	ret := funcs[puzzNumber-1](file)

	fmt.Println(ret)
}
