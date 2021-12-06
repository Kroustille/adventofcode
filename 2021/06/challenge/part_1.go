package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	input_values := strings.Split(lines[0], ",")
	values := make([]int, len(input_values))
	for i, input_value := range input_values {
		values[i] = utils.FatalReadInt(input_value)
	}
	days := 80
	for i := 0; i < days; i++ {
		for j := 0; j < len(values); j++ {
			current_value := values[j]
			if current_value == 0 {
				values[j] = 6
				values = append(values, 8)
			} else {
				values[j] = current_value - 1
			}
		}
	}

	result := len(values)
	return result, time.Since(start)
}
