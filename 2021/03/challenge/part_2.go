package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) FilterRatingsWithColumnBit(ratings []string, column int, bit string) []string {
	filtered := make([]string, 0)
	for _, r := range ratings {
		if string(r[column]) == bit {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

func (c Challenge) GetOnesCountAt(lines []string, column int) int {
	if len(lines) == 0 {
		return 0
	}
	ones_count := c.getOnesCount(lines)
	return ones_count[column]
}

func (c Challenge) getOnesCount(lines []string) []int {
	number_of_ones := make([]int, len(lines[0]))
	for _, line := range lines {
		splitted_line := strings.Split(line, "")
		for i, number := range splitted_line {
			if number == "1" {
				number_of_ones[i]++
			}
		}
	}

	return number_of_ones
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Time) {
	start := time.Now()

	i := 0
	found_oxygen_rating := ""
	found_scrubber_rating := ""
	possible_oxygen_ratings := lines
	possible_scrubber_ratings := lines

	for i < len(lines[0]) && (found_oxygen_rating == "" || found_scrubber_rating == "") {
		number_of_ones_in_oxygen := c.GetOnesCountAt(possible_oxygen_ratings, i)
		number_of_ones_in_scrubber := c.GetOnesCountAt(possible_scrubber_ratings, i)

		max_numbers_oxygen := len(possible_oxygen_ratings)
		max_number_scrubber := len(possible_scrubber_ratings)

		var oxygen_rating_bit string
		var scrubber_rating_bit string

		more_zero_than_one_oxygen := number_of_ones_in_oxygen < max_numbers_oxygen-number_of_ones_in_oxygen
		oxygen_rating_bit = "1"
		if more_zero_than_one_oxygen {
			oxygen_rating_bit = "0"
		}

		more_zero_than_one_scrubber := number_of_ones_in_scrubber < max_number_scrubber-number_of_ones_in_scrubber
		scrubber_rating_bit = "0"
		if more_zero_than_one_scrubber {
			scrubber_rating_bit = "1"
		}

		possible_oxygen_ratings = c.FilterRatingsWithColumnBit(possible_oxygen_ratings, i, oxygen_rating_bit)
		if len(possible_oxygen_ratings) == 1 && found_oxygen_rating == "" {
			found_oxygen_rating = possible_oxygen_ratings[0]
		}

		possible_scrubber_ratings = c.FilterRatingsWithColumnBit(possible_scrubber_ratings, i, scrubber_rating_bit)
		if len(possible_scrubber_ratings) == 1 && found_scrubber_rating == "" {
			found_scrubber_rating = possible_scrubber_ratings[0]
		}

		i = i + 1
	}

	final_oxygen_rating := utils.FatalReadBinary(found_oxygen_rating)
	final_scrubber_rating := utils.FatalReadBinary(found_scrubber_rating)

	result := final_oxygen_rating * final_scrubber_rating

	return result, start
}
