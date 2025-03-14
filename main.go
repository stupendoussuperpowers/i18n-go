package main

import (
	"fmt"
	"github.com/stupendoussuperpowers/i18n-go/pkg/i18n"
	"os"
	"reflect"
	"strconv"
)

type puzzleFunction func(*os.File) string

func main() {
	puzzNumber, _ := strconv.Atoi(os.Args[1])

	filepath := "tests/input%d.txt"
	if len(os.Args) > 2 && os.Args[2] == "test" {
		filepath = "tests/test-input%d.txt"
	}

	if puzzNumber <= 0 {
		fmt.Println("Puzzle number ", puzzNumber, " requires going back in time :(")
		os.Exit(1)
	}

	file, err := os.Open(fmt.Sprintf(filepath, puzzNumber))
	if err != nil {
		fmt.Printf("Input file not found. @ tests/input%d.txt", puzzNumber)
		os.Exit(1)
	}
	defer file.Close()

	puzzlesInstance := &i18n.Puzzles{}

	method := reflect.ValueOf(puzzlesInstance).MethodByName(fmt.Sprintf("Puzzle%d", puzzNumber))

	if !method.IsValid() {
		fmt.Println("Puzzle number", puzzNumber, " isn't solved yet!")
		os.Exit(1)
	}

	ret := method.Interface().(func(*os.File) string)(file)

	fmt.Println(ret)
}
