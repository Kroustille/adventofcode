package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2() {
	start := time.Now()
	_ = utils.ReadLines("input")

	result := 1
	utils.PrintResult(2, result, start)
}
