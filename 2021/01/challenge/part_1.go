package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart1() {
	start := time.Now()
	lines := utils.ReadLines("input")
	numbers := utils.ConvertLinesToIntArray(lines)

	result := c.CountMeasureIncreases(numbers)

	utils.PrintResult(1, result, start)
}
