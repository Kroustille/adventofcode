package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	c.read_index = 0
	max_magnitude := 0
	for _, line := range lines {
		for _, other_line := range lines {
			// if i == j {
			// 	break
			// }
			whole_number := c.LaunchSnailNumberParsing(line)
			new_number := c.LaunchSnailNumberParsing(other_line)

			whole_number = c.Add(whole_number, new_number)

			has_exploded := true
			has_splitted := true

			for has_exploded || has_splitted {
				has_exploded = true
				has_splitted = true
				c.last_left_value = nil
				c.right_value = nil
				whole_number, has_exploded = whole_number.Explode(&c)

				if !has_exploded {
					whole_number, has_splitted = whole_number.Split()
				}
			}

			current_magnitude := whole_number.Magnitude()
			if current_magnitude > max_magnitude {
				log.Println(whole_number)
				max_magnitude = current_magnitude
			}
		}
	}

	result := max_magnitude
	return result, time.Since(start)
}
