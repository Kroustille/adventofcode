package challenge

import "math"

type Challenge struct {
}

func (c Challenge) ComputeRequiredFuelForOneCrabPart1(crab_position, desired_position int) int {
	return int(math.Abs(float64(desired_position) - float64(crab_position)))
}

func (c Challenge) ComputeRequiredFuelForOneCrabPart2(crab_position, desired_position int) int {
	distance_between_points := c.ComputeRequiredFuelForOneCrabPart1(crab_position, desired_position)
	required_fuel := 0
	for i := 0; i <= distance_between_points; i++ {
		required_fuel += i
	}
	return required_fuel
}
