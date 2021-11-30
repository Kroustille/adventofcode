package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/2020/utils"
)

const TOTAL_RESULT = 2020

func (c Challenge) ResolvePart1() {
	start := time.Now()
	lines := utils.ReadLines("input")

	numbers := c.LinesToNumbers(lines)

	for _, line := range lines {
		number := utils.FatalReadInt(line)

		expected_result := TOTAL_RESULT - number
		if numbers[expected_result] == true {
			result := number * (TOTAL_RESULT - number)
			utils.PrintResult(1, result, start)
			return
		}
	}
}
