package challenge

import (
	"math"
	"strings"
	"time"
)

const STEP_COUNT = 10

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	combinations := make(map[string]string, 0)

	for _, line := range lines[2:] {
		splitted_line := strings.Split(line, " -> ")
		combination_key := splitted_line[0]
		new_element := splitted_line[1]
		combinations[combination_key] = new_element
	}

	elements := strings.Split(lines[0], "")

	for step := 0; step < STEP_COUNT; step++ {
		new_elements := make([]string, 0)

		for element_index := 0; element_index < len(elements)-1; element_index++ {
			new_elements = append(new_elements, elements[element_index])
			combination_key := elements[element_index] + elements[element_index+1]
			new_element, ok := combinations[combination_key]
			if ok {
				new_elements = append(new_elements, new_element)
			}
		}

		new_elements = append(new_elements, elements[len(elements)-1])

		elements = new_elements
	}

	counts := make(map[string]int, 0)

	for _, element := range elements {
		if _, ok := counts[element]; ok {
			counts[element] += 1
		} else {
			counts[element] = 1
		}
	}

	max := 0
	min := math.MaxInt
	for _, count := range counts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	result := max - min
	return result, time.Since(start)
}
