package challenge

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/Kroustille/adventofcode/utils"
)

type Challenge struct {
}

type Octopus struct {
	energy             int
	has_flashed        bool
	adjacent_octopuses []*Octopus
}

func (o Octopus) String() string {
	return fmt.Sprintf("%d", o.energy)
}

func (c Challenge) IsIndexValid(index, max_index int) bool {
	return index >= 0 && index < max_index
}

func (c Challenge) DisplayOctopusesGrid(octopuses []*Octopus) {
	line_length := math.Sqrt(float64(len(octopuses)))
	line := ""
	log.Println("##### ##### ##### #####")
	for i, octopus := range octopuses {
		line += fmt.Sprintf("%d", octopus.energy)

		if (i+1)%int(line_length) == 0 && i != 0 {
			log.Println(line)
			line = ""
		}
	}
}

func (c Challenge) BuildOctopusesArray(lines []string) []*Octopus {
	octopuses := make([][]*Octopus, len(lines[0]))

	for line_index, line := range lines {
		splitted_line := strings.Split(line, "")
		for _, character := range splitted_line {
			octopus := &Octopus{
				energy:      utils.FatalReadInt(character),
				has_flashed: false,
			}

			octopuses[line_index] = append(octopuses[line_index], octopus)
		}
	}

	log.Println(octopuses)
	max_index := len(octopuses[0])
	log.Println(max_index)
	for line_index, octopus_line := range octopuses {
		for col_index, octopus := range octopus_line {
			if c.IsIndexValid(line_index+1, max_index) {
				if c.IsIndexValid(col_index+1, max_index) {
					octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index+1][col_index+1])
				}

				if c.IsIndexValid(col_index-1, max_index) {
					octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index+1][col_index-1])
				}

				octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index+1][col_index])
			}

			if c.IsIndexValid(line_index-1, max_index) {
				if c.IsIndexValid(col_index+1, max_index) {
					octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index-1][col_index+1])
				}

				if c.IsIndexValid(col_index-1, max_index) {
					octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index-1][col_index-1])
				}

				octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index-1][col_index])
			}

			if c.IsIndexValid(col_index+1, max_index) {
				octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index][col_index+1])
			}

			if c.IsIndexValid(col_index-1, max_index) {
				octopus.adjacent_octopuses = append(octopus.adjacent_octopuses, octopuses[line_index][col_index-1])
			}
		}
	}

	final_octopuses := make([]*Octopus, 0)
	for _, octopus_line := range octopuses {
		for _, octopus := range octopus_line {
			final_octopuses = append(final_octopuses, octopus)
		}
	}

	return final_octopuses
}

func (c Challenge) Flash(octopus *Octopus) {
	octopus.has_flashed = true

	for _, adjacent_octopus := range octopus.adjacent_octopuses {
		adjacent_octopus.energy++
		c.CheckFlash(adjacent_octopus)
	}
}

func (c Challenge) CheckFlash(octopus *Octopus) {
	if octopus.energy > 9 && !octopus.has_flashed {
		c.Flash(octopus)
	}
}
