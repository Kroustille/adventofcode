package challenge

import (
	"log"
	"time"
)

const STEP_COUNT_2 = 500

func (c Challenge) RunStep2(octopuses []*Octopus) bool {
	for _, octopus := range octopuses {
		octopus.energy++
	}

	for _, octopus := range octopuses {
		c.CheckFlash(octopus)
	}

	flash_count := 0
	for _, octopus := range octopuses {
		if octopus.has_flashed {
			flash_count += 1
			octopus.energy = 0
			octopus.has_flashed = false
		}
	}

	// log.Println(flash_count, len(octopuses))
	if flash_count == len(octopuses) {
		return true
	}

	return false
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()
	octopuses := c.BuildOctopusesArray(lines)
	log.Println("INITIAL DATA")
	c.DisplayOctopusesGrid(octopuses)
	winning_step := 0
	for i := 0; i < STEP_COUNT_2; i++ {
		all_flashed := c.RunStep2(octopuses)
		if all_flashed {
			winning_step = i + 1
			log.Println("all_flashed", all_flashed, winning_step)
			break
		}
		log.Println("STEP ", i+1)
		c.DisplayOctopusesGrid(octopuses)
	}

	result := winning_step
	return result, time.Since(start)
}
