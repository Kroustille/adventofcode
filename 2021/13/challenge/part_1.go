package challenge

import (
	"fmt"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) Key() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Challenge) FoldAlongX(all_coordinates map[string]Coordinates, fold_coordinate int) map[string]Coordinates {
	new_coordinates := make(map[string]Coordinates, 0)
	for _, coordinate := range all_coordinates {
		if coordinate.x < fold_coordinate {
			new_coordinates[coordinate.Key()] = coordinate
		} else {
			folded_coordinate := Coordinates{
				x: fold_coordinate - (coordinate.x - fold_coordinate),
				y: coordinate.y,
			}
			new_coordinates[folded_coordinate.Key()] = folded_coordinate
		}
	}

	return new_coordinates
}

func (c Challenge) FoldAlongY(all_coordinates map[string]Coordinates, fold_coordinate int) map[string]Coordinates {
	new_coordinates := make(map[string]Coordinates, 0)
	for _, coordinate := range all_coordinates {
		if coordinate.y < fold_coordinate {
			new_coordinates[coordinate.Key()] = coordinate
		} else {
			folded_coordinate := Coordinates{
				x: coordinate.x,
				y: fold_coordinate - (coordinate.y - fold_coordinate),
			}
			new_coordinates[folded_coordinate.Key()] = folded_coordinate
		}
	}

	return new_coordinates
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	all_coordinates := make(map[string]Coordinates, 0)
	fold_start_index := 0
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

		all_coordinates[new_coordinates.Key()] = new_coordinates
	}

	splitted_line := strings.Split(lines[fold_start_index], " ")
	fold_instructions := strings.Split(splitted_line[2], "=")
	coordinate_type := fold_instructions[0]
	coordinate_value := utils.FatalReadInt(fold_instructions[1])
	if coordinate_type == "x" {
		all_coordinates = c.FoldAlongX(all_coordinates, coordinate_value)
	}

	if coordinate_type == "y" {
		all_coordinates = c.FoldAlongY(all_coordinates, coordinate_value)
	}

	result := len(all_coordinates)

	return result, time.Since(start)
}
