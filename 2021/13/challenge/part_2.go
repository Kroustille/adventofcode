package challenge

import (
	"log"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	all_coordinates := make(map[string]Coordinates, 0)
	fold_start_index := 0
	max_x := 0
	max_y := 0
	for i, line := range lines {
		if line == "" {
			fold_start_index = i + 1
			break
		}
		splitted_line := strings.Split(line, ",")
		new_coordinates := Coordinates{
			x: utils.FatalReadInt(splitted_line[0]),
			y: utils.FatalReadInt(splitted_line[1]),
		}

		if new_coordinates.x > max_x {
			max_x = new_coordinates.x
		}

		if new_coordinates.y > max_y {
			max_y = new_coordinates.y
		}

		all_coordinates[new_coordinates.Key()] = new_coordinates
	}

	for i := fold_start_index; i < len(lines); i++ {
		log.Println(lines[i])
		splitted_line := strings.Split(lines[i], " ")
		fold_instructions := strings.Split(splitted_line[2], "=")
		coordinate_type := fold_instructions[0]
		coordinate_value := utils.FatalReadInt(fold_instructions[1])
		if coordinate_type == "x" {
			all_coordinates = c.FoldAlongX(all_coordinates, coordinate_value)
			max_x /= 2
		}

		if coordinate_type == "y" {
			all_coordinates = c.FoldAlongY(all_coordinates, coordinate_value)
			max_y /= 2
		}
	}

	for y := 0; y < max_y; y++ {
		display_line := ""
		for x := 0; x < max_x; x++ {
			coordinate := Coordinates{
				x: x,
				y: y,
			}
			if _, ok := all_coordinates[coordinate.Key()]; ok {
				display_line += "#"
			} else {
				display_line += "."
			}
		}
		log.Println(display_line)
	}

	result := len(all_coordinates)

	return result, time.Since(start)
}
