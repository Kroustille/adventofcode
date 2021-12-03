package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart1() {
	start := time.Now()
	lines := utils.ReadLines("input")

	depth := 0
	horizontal := 0
	for _, line := range lines {
		data := strings.Split(line, " ")
		move := utils.FatalReadInt(data[1])
		switch data[0] {
		case FORWARD:
			horizontal += move
			break
		case UP:
			depth -= move
			break
		case DOWN:
			depth += move
			break
		}
	}

	result := depth * horizontal
	utils.PrintResult(1, result, start)
}
