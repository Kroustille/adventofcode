package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart1() {
	start := time.Now()
	lines := utils.ReadLines("input")

	max_numbers := len(lines)
	number_of_ones := make([]int, len(lines[0]))
	for _, line := range lines {
		splitted_line := strings.Split(line, "")
		for i, number := range splitted_line {
			if number == "1" {
				number_of_ones[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, column_ones := range number_of_ones {
		if column_ones > max_numbers/2 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	gamma_number := utils.FatalReadBinary(gamma)
	epsilon_number := utils.FatalReadBinary(epsilon)

	result := gamma_number * epsilon_number
	utils.PrintResult(1, result, start)
}
