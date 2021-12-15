package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) IsIndexValid(index, max_value int) bool {
	return index >= 0 && index < max_value
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	lowest_points := make([]int, 0)

	line_max_value := len(lines)
	character_max_value := len(lines[0])
	for line_index, line := range lines {
		current_splitted_line := strings.Split(line, "")

		for character_index, current_character := range current_splitted_line {
			current_number := utils.FatalReadInt(current_character)
			is_lowest := true

			if c.IsIndexValid(line_index-1, line_max_value) {
				splitted_other_line := strings.Split(lines[line_index-1], "")
				other_number := splitted_other_line[character_index]
				is_lowest = is_lowest && current_number < utils.FatalReadInt(other_number)
			}

			if c.IsIndexValid(line_index+1, line_max_value) {
				splitted_other_line := strings.Split(lines[line_index+1], "")
				other_number := splitted_other_line[character_index]
				is_lowest = is_lowest && current_number < utils.FatalReadInt(other_number)
			}

			if c.IsIndexValid(character_index-1, character_max_value) {
				other_number := line[character_index-1]
				is_lowest = is_lowest && current_number < utils.FatalReadInt(string(other_number))
			}

			if c.IsIndexValid(character_index+1, character_max_value) {
				other_number := line[character_index+1]
				is_lowest = is_lowest && current_number < utils.FatalReadInt(string(other_number))
			}

			if is_lowest {
				lowest_points = append(lowest_points, current_number)
			}
		}
	}

	score := 0
	for _, points := range lowest_points {
		score += points + 1
	}

	result := score
	return result, time.Since(start)
}
