package i18n

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func (p *Puzzles) Puzzle3(file *os.File) string {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		l := scanner.Text()

		runes := utf8.RuneCountInString(l)
		if runes > 12 || runes < 4 {
			continue
		}

		var digit bool = false
		var upper bool = false
		var lower bool = false
		var ascii bool = false

		for _, c := range l {
			if unicode.IsDigit(c) {
				digit = true
			}

			if unicode.IsUpper(c) {
				upper = true
			}

			if unicode.IsLower(c) {
				lower = true
			}

			if c > unicode.MaxASCII {
				ascii = true
			}

			if lower && digit && upper && ascii {
				count += 1
				break
			}
		}

	}

	return strconv.Itoa(count)
}
