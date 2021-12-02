package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2() {
	start := time.Now()
	lines := utils.ReadLines("input")

	aim := 0
	depth := 0
	horizontal := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		move := utils.FatalReadInt(data[1])
		switch data[0] {
		case FORWARD:
			horizontal += move
			depth += aim * move
			break
		case UP:
			aim -= move
			break
		case DOWN:
			aim += move
			break
		}
	}

	result := horizontal * depth
	utils.PrintResult(2, result, start)
}
