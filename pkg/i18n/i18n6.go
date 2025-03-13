package i18n

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func unMismatch(input string) string {
	encoder := charmap.ISO8859_1.NewEncoder()
	latinBytes := []byte(input)
	ans, err := encoder.Bytes(latinBytes)

	if err != nil {
		return ""
	}

	return string(ans)
}

func extract(input string) (int, int, byte) {
	pos := 0
	runes := utf8.RuneCountInString(input)

	for pos = 0; pos < runes; pos++ {
		if input[pos] != '.' {
			break
		}
	}

	return runes, pos, input[pos]
}

func match(input string, word string) bool {

	runes, pos, ch := extract(input)

	wordRunes := []rune(word)

	if utf8.RuneCountInString(word) != runes {
		return false
	}

	if wordRunes[pos] != rune(ch) {
		return false
	}

	return true
}

func PuzzleSix(file *os.File) string {

	scanner := bufio.NewScanner(file)

	wordList := []string{}
	crossWord := []string{}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i = i + 1

		if line == "" {
			break
		}

		if i%3 == 0 {
			line = unMismatch(line)
		}

		if i%5 == 0 {
			line = unMismatch(line)
		}
		wordList = append(wordList, line)

		fmt.Println(line)
	}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		crossWord = append(crossWord, strings.TrimSpace(line))
	}

	ans := 0

	for _, word := range crossWord {
		for i := 0; i < len(wordList); i++ {
			if match(word, wordList[i]) {
				ans += i + 1
				fmt.Println(wordList[i])
				fmt.Println(word)
				break
			}
		}
	}

	return strconv.Itoa(ans)
}
