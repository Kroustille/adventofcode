package challenge

import (
	"fmt"
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	wind_values := make(map[string]int, 0)
	for _, line := range lines {
		wind_vector := c.ParseLine(line)
		is_valid := wind_vector.IsVerticalOrHorizontal()
		if !is_valid {
			continue
		}

		loop_start_x := wind_vector.start.x
		loop_end_x := wind_vector.end.x
		if wind_vector.end.x < wind_vector.start.x {
			loop_start_x = wind_vector.end.x
			loop_end_x = wind_vector.start.x
		}
		loop_start_y := wind_vector.start.y
		loop_end_y := wind_vector.end.y
		if wind_vector.end.y < wind_vector.start.y {
			loop_start_y = wind_vector.end.y
			loop_end_y = wind_vector.start.y
		}
		for i := loop_start_x; i <= loop_end_x; i++ {
			for j := loop_start_y; j <= loop_end_y; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				_, ok := wind_values[key]
				if !ok {
					wind_values[key] = 1
				} else {
					wind_values[key] += 1
				}
			}
		}
	}
	result := 0
	for _, wind_value := range wind_values {
		if wind_value >= 2 {
			result += 1
		}
	}

	return result, time.Since(start)
}
