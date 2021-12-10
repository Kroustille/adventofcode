package challenge

import (
	"sort"
	"time"
)

func (c Challenge) FindLastOpeningCharacter(line string, index int, last_opening_character OpeningCharacter) OpeningCharacter {
	if index == len(line) {
		return last_opening_character
	}

	character := string(line[index])
	if c.IsOpeningCharacter(character) {
		new_opening_character := OpeningCharacter{
			current:                character,
			last_opening_character: &last_opening_character,
		}
		return c.FindLastOpeningCharacter(line, index+1, new_opening_character)
	}

	return c.FindLastOpeningCharacter(line, index+1, *last_opening_character.last_opening_character)
}

func (c Challenge) TransformOpeningToClosingCharacter(opening_character string) string {
	switch opening_character {
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	case "(":
		return ")"
	}

	return ""
}

func (c Challenge) BuildSequenceFromOpeningCharacters(last_opening_character OpeningCharacter) []string {
	complete_sequence := make([]string, 0)
	for last_opening_character.current != "" {
		closing_character := c.TransformOpeningToClosingCharacter(last_opening_character.current)
		complete_sequence = append(complete_sequence, closing_character)
		last_opening_character = *last_opening_character.last_opening_character
	}

	return complete_sequence
}

func (c Challenge) ComputeLineScore(line string) int {
	first_opening_character := OpeningCharacter{
		current: "",
	}
	is_corrupted := c.FindCorruptedCharacter(line, 0, first_opening_character) != ""
	if is_corrupted {
		return 0
	}

	last_opening_character := c.FindLastOpeningCharacter(line, 0, first_opening_character)
	complete_sequence := c.BuildSequenceFromOpeningCharacters(last_opening_character)
	final_score := 0
	for _, character := range complete_sequence {
		final_score *= 5

		switch character {
		case ")":
			final_score += 1
			break
		case "]":
			final_score += 2
			break
		case "}":
			final_score += 3
			break
		case ">":
			final_score += 4
			break
		}
	}

	return final_score
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	final_scores := make([]int, 0)
	for _, line := range lines {
		line_score := c.ComputeLineScore(line)
		if line_score != 0 {
			final_scores = append(final_scores, line_score)
		}
	}

	sort.Ints(final_scores[:])
	result := final_scores[len(final_scores)/2]
	return result, time.Since(start)
}
