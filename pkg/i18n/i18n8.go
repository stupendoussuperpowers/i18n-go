package i18n

import (
	"bufio"
	// "fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	// "strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func isVowelAccented(s string) (int, bool, bool, bool) {
	de := norm.NFD.String(s)

	vowel := false
	accented := false
	digit := false
	number := 0

	vowOrds := []int{97, 101, 105, 111, 117, 65, 69, 73, 79, 85}

	for _idx, dr := range de {
		if _idx == 0 {
			number = int(dr)
		}

		if unicode.Is(unicode.Mn, dr) {
			accented = true
		}

		if slices.Contains(vowOrds, int(dr)) {
			vowel = true
		}

		if int(dr) <= 57 && int(dr) >= 48 {
			digit = true
		}
	}

	return number, digit, vowel, accented
}

func (p *Puzzles) Puzzle8(file *os.File) string {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		actualLength := utf8.RuneCountInString(line)

		if actualLength > 12 || actualLength < 4 {
			continue
		}

		accented := false
		digit := false
		vowel := false
		consonant := false

		letters := make(map[int]bool)

		for _, ch := range []rune(line) {
			c := strings.ToLower(string(ch))
			l, d, v, a := isVowelAccented(c)

			if letters[l] {
				digit = false
				break
			} else {
				letters[l] = true
			}

			accented = accented || a
			vowel = vowel || v
			digit = digit || d
			consonant = consonant || !v
		}

		if accented && vowel && digit && consonant {
			count += 1
		}
	}

	return strconv.Itoa(count)
}
