package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2() {
	start := time.Now()
	lines := utils.ReadLines("input")
	numbers := utils.ConvertLinesToIntArray(lines)
	measure_windows := make([]int, len(numbers))
	for i := 0; i < len(numbers)-2; i++ {
		measure_windows[i] = numbers[i] + numbers[i+1] + numbers[i+2]
	}

	result := c.CountMeasureIncreases(measure_windows)

	utils.PrintResult(2, result, start)
}
