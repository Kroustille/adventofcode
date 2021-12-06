package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	input_values := strings.Split(lines[0], ",")

	fishes := make(map[int]int, 9)
	for i := 0; i <= 8; i++ {
		fishes[i] = 0
	}

	for _, input_value := range input_values {
		key := utils.FatalReadInt(input_value)
		fishes[key]++
	}

	days := 256
	for i := 0; i < days; i++ {
		new_fishes := make(map[int]int, 9)
		for fish_day := 8; fish_day >= 0; fish_day-- {
			if fish_day == 0 {
				new_fishes[8] += fishes[0]
				new_fishes[6] += fishes[0]
			} else {
				new_fishes[fish_day-1] = fishes[fish_day]
			}
		}
		fishes = new_fishes
	}

	result := 0
	for i := 0; i <= 8; i++ {
		result += fishes[i]
	}

	return result, time.Since(start)
}
