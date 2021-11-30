package challenge

import (
	"log"
	"strconv"
	"time"

	"github.com/Kroustille/adventofcode/2020/utils"
)

func (c Challenge) ResolvePart2() {
	start := time.Now()
	lines := utils.ReadLines("input")

	numbers := make(map[int]bool, len(lines))
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		numbers[number] = true
	}

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		for _, second_line := range lines {
			second_number := utils.FatalReadInt(second_line)
			expected_result := TOTAL_RESULT - number - second_number
			if numbers[expected_result] == true {
				result := number * second_number * (TOTAL_RESULT - number - second_number)
				utils.PrintResult(2, result, start)
				return
			}
		}
	}
}
