package challenge

import (
	"strings"
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	sum_of_digits := 0
	for _, line := range lines {
		splitted_line := strings.Split(line, " | ")
		output := splitted_line[1]

		digits := strings.Split(output, " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				sum_of_digits++
				break
			case 4:
				sum_of_digits++
				break
			case 3:
				sum_of_digits++
				break
			case 7:
				sum_of_digits++
				break
			default:
				break
			}
		}
	}

	result := sum_of_digits
	return result, time.Since(start)
}
