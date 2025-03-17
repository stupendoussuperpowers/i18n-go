package i18n

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (p *Puzzles) Puzzle7(file *os.File) string {
	scanner := bufio.NewScanner(file)

	i := 1
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		rawTime, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			log.Fatal(err)
			return "Hm."
		}

		subMin, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
			return "Hm."
		}

		addMin, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
			return "Hm."
		}

		hfx, err := time.LoadLocation("America/Halifax")
		if err != nil {
			log.Fatal(err)
			return ""
		}

		san, err := time.LoadLocation("America/Santiago")
		if err != nil {
			log.Fatal(err)
			return ""
		}

		var loc *time.Location = san

		rawTime = rawTime.In(loc)

		if rawTime.In(hfx).UTC().Format("-07:00") == rawTime.Format("-07:00") {
			loc = hfx
		}

		rawTime = rawTime.Add(time.Duration(-subMin) * time.Minute)
		rawTime = rawTime.Add(time.Duration(addMin) * time.Minute)
		rawTime = rawTime.In(loc)

		fmt.Println(loc, rawTime, rawTime.Hour(), count)
		count += rawTime.Hour() * i
		i += 1
	}

	return strconv.Itoa(count)
}
