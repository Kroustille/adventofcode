package challenge

import (
	"log"
	"time"
)

const STEP_COUNT = 100

func (c Challenge) RunStep1(octopuses []*Octopus) int {
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

	return flash_count
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()
	octopuses := c.BuildOctopusesArray(lines)
	log.Println("INITIAL DATA")
	c.DisplayOctopusesGrid(octopuses)
	total_flash_count := 0
	for i := 0; i < STEP_COUNT; i++ {
		total_flash_count += c.RunStep1(octopuses)
		log.Println("STEP ", i+1)
		c.DisplayOctopusesGrid(octopuses)
	}

	result := total_flash_count
	return result, time.Since(start)
}
