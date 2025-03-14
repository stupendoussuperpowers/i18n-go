package i18n

import (
	"bufio"
	"os"
	"strconv"
)

func (p *Puzzles) Puzzle5(file *os.File) string {
	scanner := bufio.NewScanner(file)

	start := 0
	count := 0

	for scanner.Scan() {
		runes := []rune(scanner.Text())

		linecount := len(runes)

		if runes[start] == '\U0001F4A9' {
			count += 1
		}

		start = (start + 2) % linecount
	}

	return strconv.Itoa(count)
}
