package challenge

import (
	"log"
	"strings"
	"unicode"
)

type Challenge struct {
}

type Cave struct {
	code           string
	already_viewed bool
	linked_caves   []*Cave
}

func (c Cave) String() string {
	return c.code
}

func (c Challenge) BuildCaves(lines []string) *Cave {
	all_caves := make(map[string]*Cave, 0)

	for _, line := range lines {
		splitted_line := strings.Split(line, "-")
		cave_one_code := splitted_line[0]
		cave_two_code := splitted_line[1]

		cave_one := all_caves[cave_one_code]
		cave_two := all_caves[cave_two_code]

		if cave_one != nil {
			if cave_two != nil {
				cave_one.linked_caves = append(cave_one.linked_caves, cave_two)
				cave_two.linked_caves = append(cave_two.linked_caves, cave_one)
			} else {
				cave_two = &Cave{
					code:         cave_two_code,
					linked_caves: []*Cave{cave_one},
				}
				all_caves[cave_two_code] = cave_two
				cave_one.linked_caves = append(cave_one.linked_caves, cave_two)
			}
		} else {
			if cave_two != nil {
				cave_one = &Cave{
					code:         cave_one_code,
					linked_caves: []*Cave{cave_two},
				}
				all_caves[cave_one_code] = cave_one
				cave_two.linked_caves = append(cave_two.linked_caves, cave_one)
			} else {
				cave_one = &Cave{
					code:         cave_one_code,
					linked_caves: []*Cave{},
				}
				cave_two = &Cave{
					code:         cave_two_code,
					linked_caves: []*Cave{cave_one},
				}
				cave_one.linked_caves = append(cave_one.linked_caves, cave_two)
				all_caves[cave_one_code] = cave_one
				all_caves[cave_two_code] = cave_two
			}
		}
	}

	var start_cave *Cave
	for _, cave := range all_caves {
		if cave.IsStart() {
			start_cave = cave
			break
		}
	}

	if start_cave == nil {
		log.Fatal("start cave not found")
	}

	return start_cave
}

func (c Cave) IsSmall() bool {
	for _, r := range string(c.code[0]) {
		if unicode.IsLower(r) {
			return true
		}
	}

	return false
}

func (c Cave) IsStart() bool {
	return c.code == "start"
}

func (c Cave) IsEnd() bool {
	return c.code == "end"
}
