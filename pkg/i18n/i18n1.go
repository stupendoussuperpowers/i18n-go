package i18n

import (
	"bufio"
	"os"
	"strconv"
	"unicode/utf8"
)

func (p *Puzzles) Puzzle1(file *os.File) string {
	finans := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		ans := 0

		length := len(line)
		runes := utf8.RuneCountInString(line)

		if length <= 160 {
			ans += 11
		}

		if runes <= 140 {
			ans = min(ans+7, 13)
		}

		finans += ans
	}

	return strconv.Itoa(finans)
}
