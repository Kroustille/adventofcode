package challenge

import (
	"time"
)

func (c Challenge) GetScore(character string) int {
	switch character {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	first_opening_character := OpeningCharacter{
		current: "",
	}

	final_score := 0
	for _, line := range lines {
		first_corrupted_character := c.FindCorruptedCharacter(line, 0, first_opening_character)
		if first_corrupted_character != "" {
			character_score := c.GetScore(first_corrupted_character)
			final_score += character_score
		}
	}

	result := final_score
	return result, time.Since(start)
}
