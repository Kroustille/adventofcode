package challenge

import (
	"math"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	first_line := lines[0]
	splitted_line := strings.Split(first_line, ",")

	crab_positions := make([]int, len(splitted_line))
	max_position := 0
	for i, line := range splitted_line {
		crab_position := utils.FatalReadInt(line)
		crab_positions[i] = crab_position

		if crab_position > max_position {
			max_position = crab_position
		}
	}

	minimal_required_fuel := math.MaxInt
	for desired_position := 0; desired_position <= max_position; desired_position++ {
		required_fuel_for_position := 0
		for _, crab_position := range crab_positions {
			required_fuel_for_position += c.ComputeRequiredFuelForOneCrabPart2(crab_position, desired_position)
		}

		if required_fuel_for_position < minimal_required_fuel {
			minimal_required_fuel = required_fuel_for_position
		}
	}

	result := minimal_required_fuel
	return result, time.Since(start)
}
