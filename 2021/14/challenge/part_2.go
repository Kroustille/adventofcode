package challenge

import (
	"log"
	"math"
	"strings"
	"time"
)

const STEP_COUNT_2 = 40

func DisplayCombinations(combinations map[string]int) {
	log.Println("#### COMBINATIONS ####")
	for _, combination := range combinations {
		log.Println(combination)
	}
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	transformations := make(map[string]string, 0)
	combinations_count := make(map[string]int, 0)

	for _, line := range lines[2:] {
		splitted_line := strings.Split(line, " -> ")
		combination_key := splitted_line[0]
		new_element := splitted_line[1]
		combinations_count[combination_key] = 0
		transformations[combination_key] = new_element
	}

	first_elements := strings.Split(lines[0], "")
	for element_index := 0; element_index < len(first_elements)-1; element_index++ {
		combination_key := first_elements[element_index] + first_elements[element_index+1]
		combinations_count[combination_key]++
	}

	for i := 0; i < STEP_COUNT_2; i++ {
		new_combinations_count := make(map[string]int, len(combinations_count))
		for key, count := range combinations_count {
			transform_element := transformations[key]
			splitted_key := strings.Split(key, "")
			new_combination_key_left := splitted_key[0] + transform_element
			new_combination_key_right := transform_element + splitted_key[1]

			new_combinations_count[new_combination_key_left] += count
			new_combinations_count[new_combination_key_right] += count
		}

		combinations_count = new_combinations_count
	}

	elements_count := make(map[string]int, 0)
	for key, count := range combinations_count {
		splitted_key := strings.Split(key, "")
		second_element := splitted_key[1]

		elements_count[second_element] += count
	}

	elements_count[first_elements[0]]++

	log.Println(elements_count)

	max := 0
	min := math.MaxInt
	for _, count := range elements_count {
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
