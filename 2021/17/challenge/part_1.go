package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	goal := c.ParseGoal(lines[0])

	final_highest_point := 0
	for x_velocity := 0; x_velocity < 10000; x_velocity++ {
		for y_velocity := 0; y_velocity < 10000; y_velocity++ {
			probe := Probe{
				x:          0,
				y:          0,
				x_velocity: x_velocity,
				y_velocity: y_velocity,
			}
			highest_point := 0
			for !goal.IsProbeInGoal(probe) {
				probe.Step()
				if probe.y > highest_point {
					highest_point = probe.y
				}

				if goal.IsProbeTooFar(probe) {
					break
				}
			}

			if goal.IsProbeInGoal(probe) {
				log.Println("probe-in-goal:", "x_velocity =", x_velocity, "y_velocity =", y_velocity, "highest_point", highest_point)

				if highest_point > final_highest_point {
					log.Println("new-highest-point")
					final_highest_point = highest_point
				}
			}
		}
	}

	result := final_highest_point
	return result, time.Since(start)
}
