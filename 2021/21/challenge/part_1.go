package challenge

import (
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
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

	dice := Dice{
		current_value: 0,
		max_value:     100,
	}

	first_player_turn := true
	for first_player.score < 1000 && second_player.score < 1000 {
		distance := 0
		for i := 0; i < 3; i++ {
			distance += dice.Roll()
		}

		if first_player_turn {
			first_player.Move(distance)

			first_player_turn = false
		} else {
			second_player.Move(distance)

			first_player_turn = true
		}
	}

	loser_score := first_player.score
	if first_player.score >= 1000 {
		loser_score = second_player.score
	}

	result := dice.times_rolled * loser_score
	return result, time.Since(start)
}
