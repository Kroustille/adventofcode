package challenge

type Challenge struct {
	first_player            *Player
	second_player           *Player
	first_player_universes  *int
	second_player_universes *int
	possible_values         map[int]int
}

func (c *Challenge) InitPossibleValues() {
	c.possible_values = make(map[int]int, 0)

	for dice_1_value := 1; dice_1_value <= 3; dice_1_value++ {
		for dice_2_value := 1; dice_2_value <= 3; dice_2_value++ {
			for dice_3_value := 1; dice_3_value <= 3; dice_3_value++ {
				c.possible_values[dice_1_value+dice_2_value+dice_3_value]++
			}
		}
	}
}

type Player struct {
	position int
	score    int
}

func (player Player) MoveValue(distance int) Player {
	new_position := ((player.position + distance - 1) % 10) + 1
	return Player{
		position: new_position,
		score:    player.score + new_position,
	}
}

func (player *Player) Move(distance int) {
	player.position = ((player.position + distance - 1) % 10) + 1
	player.score += player.position
}

type Dice struct {
	current_value int
	times_rolled  int
	max_value     int
}

func (dice *Dice) Roll() int {
	new_value := dice.current_value + 1
	if new_value > dice.max_value {
		new_value = 1
	}

	dice.current_value = new_value
	dice.times_rolled++

	return dice.current_value
}
