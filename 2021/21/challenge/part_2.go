package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) Step(first_player, second_player Player, first_player_turn bool, count int) {
	if first_player.score >= 21 {
		*c.first_player_universes += count
		return
	} else if second_player.score >= 21 {
		*c.second_player_universes += count
		return
	}

	for dice_value, dice_count := range c.possible_values {
		if first_player_turn {
			c.Step(first_player.MoveValue(dice_value), second_player, !first_player_turn, dice_count*count)
		} else {
			c.Step(first_player, second_player.MoveValue(dice_value), !first_player_turn, dice_count*count)
		}
	}
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	first_player_position := utils.FatalReadInt(lines[0])
	second_player_position := utils.FatalReadInt(lines[1])

	first_player := Player{
		position: first_player_position,
		score:    0,
	}

	second_player := Player{
		position: second_player_position,
		score:    0,
	}

	first_player_universes := 0
	second_player_universes := 0

	c.first_player_universes = &first_player_universes
	c.second_player_universes = &second_player_universes
	c.InitPossibleValues()

	first_player_turn := true
	c.Step(first_player, second_player, first_player_turn, 1)

	result := *c.first_player_universes

	if *c.second_player_universes > *c.first_player_universes {
		result = *c.second_player_universes
	}
	return result, time.Since(start)
}
