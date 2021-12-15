package main

import (
	"github.com/Kroustille/adventofcode/2021/12/challenge"
	"github.com/Kroustille/adventofcode/utils"
)

func main() {
	c := challenge.Challenge{}

	lines := utils.ReadLines("input")

	// result_1, time_1 := c.ResolvePart1(lines)
	// utils.PrintResultDuration(1, result_1, time_1)

	result_2, time_2 := c.ResolvePart2(lines)
	utils.PrintResultDuration(2, result_2, time_2)
}
