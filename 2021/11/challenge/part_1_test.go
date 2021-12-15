package challenge

import (
	"log"
	"strings"
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func checkOctopusesArray(t *testing.T, octopuses []*Octopus, input_path string) {
	expected_energy_one_step_lines := utils.ReadLines(input_path)

	expected_energy_levels := make([]int, 0)
	for _, line := range expected_energy_one_step_lines {
		splitted_line := strings.Split(line, "")
		for _, character := range splitted_line {
			expected_energy_levels = append(expected_energy_levels, utils.FatalReadInt(character))
		}
	}

	octopuses_energy_levels := make([]int, len(octopuses))
	for i, octopus := range octopuses {
		octopuses_energy_levels[i] = octopus.energy
	}

	assert.Equal(t, expected_energy_levels, octopuses_energy_levels)
}

func TestRunStepOneTime(t *testing.T) {
	lines := utils.ReadLines("../test_input")
	c := Challenge{}

	octopuses := c.BuildOctopusesArray(lines)
	_ = c.RunStep1(octopuses)

	checkOctopusesArray(t, octopuses, "../one_step_energy")
	// assert.Equal(t, 0, flash_count)
}

func TestRunStepTwoTimes(t *testing.T) {
	lines := utils.ReadLines("../test_input")
	c := Challenge{}

	octopuses := c.BuildOctopusesArray(lines)
	_ = c.RunStep1(octopuses)
	log.Println(octopuses)
	_ = c.RunStep1(octopuses)

	checkOctopusesArray(t, octopuses, "../two_step_energy")
	// assert.Equal(t, 35, flash_count)
}
