package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	c.read_index = 0
	whole_number := c.LaunchSnailNumberParsing(lines[0])
	log.Println("start number  ", whole_number)
	for _, line := range lines[1:] {
		new_number := c.LaunchSnailNumberParsing(line)
		whole_number = c.Add(whole_number, new_number)
		log.Println("after addition", whole_number)

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
				if has_splitted {
					log.Println("after split   ", whole_number)
				}
			} else {
				log.Println("after explode ", whole_number)
			}
		}
	}

	result := whole_number.Magnitude()
	return result, time.Since(start)
}
