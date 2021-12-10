package challenge

import "strings"

type Challenge struct {
}

func (c Challenge) FindCorruptedCharacter(line string, index int, opening_character OpeningCharacter) string {
	if index == len(line) {
		return ""
	}

	character := string(line[index])
	if c.IsOpeningCharacter(string(character)) {
		new_opening_character := OpeningCharacter{
			last_opening_character: &opening_character,
			current:                character,
		}
		return c.FindCorruptedCharacter(line, index+1, new_opening_character)
	}

	if c.IsCorruptedClosingCharacter(opening_character.current, character) {
		return character
	}

	return c.FindCorruptedCharacter(line, index+1, *opening_character.last_opening_character)
}

type OpeningCharacter struct {
	last_opening_character *OpeningCharacter
	current                string
}

func (c Challenge) IsOpeningCharacter(character string) bool {
	opening_characters := "{<(["
	return strings.Contains(opening_characters, character)
}

func (c Challenge) IsCorruptedClosingCharacter(opening_character, closing_character string) bool {
	switch opening_character {
	case "{":
		return closing_character != "}"
	case "[":
		return closing_character != "]"
	case "<":
		return closing_character != ">"
	case "(":
		return closing_character != ")"
	}

	return true
}
