package i18n

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func PuzzleTwo(file *os.File) string {

	recordings := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		to, err := time.Parse(time.RFC3339, scanner.Text())

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		recordings[to.UTC().String()] += 1

		if recordings[to.UTC().String()] >= 4 {
			return to.UTC().String()
		}
	}

	return ""
}
