package challenge

import (
	"log"
	"strconv"
)

type Challenge struct {
}

func (c Challenge) LinesToNumbers(lines []string) map[int]bool {
	numbers := make(map[int]bool, len(lines))
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		numbers[number] = true
	}
	return numbers
}
