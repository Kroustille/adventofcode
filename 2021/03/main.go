package main

import (
	"github.com/Kroustille/adventofcode/2021/03/challenge"
	"github.com/Kroustille/adventofcode/utils"
)

func main() {
	c := challenge.Challenge{}
	lines := utils.ReadLines("input")

	// c.ResolvePart1()

	result_2, part_2_time := c.ResolvePart2(lines)
	utils.PrintResult(2, result_2, part_2_time)
}
