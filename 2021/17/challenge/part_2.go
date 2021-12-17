package challenge

import (
	"time"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	goal := c.ParseGoal(lines[0])

	winning_velocities_count := 0
	for x_velocity := 0; x_velocity < 10000; x_velocity++ {
		for y_velocity := -10000; y_velocity < 10000; y_velocity++ {
			probe := Probe{
				x:          0,
				y:          0,
				x_velocity: x_velocity,
				y_velocity: y_velocity,
			}
			for !goal.IsProbeInGoal(probe) {
				probe.Step()

				if goal.IsProbeTooFar(probe) {
					break
				}
			}

			if goal.IsProbeInGoal(probe) {
				winning_velocities_count++
			}
		}
	}

	result := winning_velocities_count
	return result, time.Since(start)
}
