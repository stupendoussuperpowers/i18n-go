package i18n

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func sanerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 1, data[:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func returnMinDiff(tz string, tm string, tz2 string, tm2 string) float64 {

	location, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	layout := "Jan 02, 2006, 15:04"

	t, err := time.ParseInLocation(layout, tm, location)

	if err != nil {
		fmt.Println("Error", err, tm)
		os.Exit(1)
	}

	location, err = time.LoadLocation(tz2)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	u, err := time.ParseInLocation(layout, tm2, location)

	if err != nil {
		fmt.Println("Error", err, tm2)

		os.Exit(1)
	}

	return u.Sub(t).Minutes()

}

func PuzzleFour(file *os.File) string {

	travelt := 0.0

	scanner := bufio.NewScanner(file)
	scanner.Split(sanerSplit)

	for scanner.Scan() {
		ticket := strings.TrimSpace(scanner.Text())

		tzFrom := strings.TrimSpace(ticket[11:41])
		tmFrom := strings.TrimSpace(ticket[42:61])

		tzTo := strings.TrimSpace(ticket[73:103])
		tmTo := strings.TrimSpace(ticket[104:])

		travelt += returnMinDiff(tzFrom, tmFrom, tzTo, tmTo)

	}

	return strconv.FormatFloat(travelt, 'f', -1, 64)
}
