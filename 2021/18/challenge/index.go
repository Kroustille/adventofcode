package challenge

import (
	"log"

	"github.com/Kroustille/adventofcode/utils"
)

const (
	START_SNAIL_NUMBER_CHARACTER = "["
	END_SNAIL_NUMBER_CHARACTER   = "]"
	SEPARATION_CHARACTER         = ","
	INITIAL_DEPTH                = 1
)

type SnailNumber interface {
	String() string
	Magnitude() int
	IncreaseDepth()
	Explode(*Challenge) (SnailNumber, bool)
	Split() (SnailNumber, bool)
	AddRightValue(int)
	AddLeftValue(int)
}

type Challenge struct {
	read_index      int
	line            string
	current_depth   int
	last_left_value *Value
	right_value     *Value
}

func (c Challenge) Add(left_number, right_number SnailNumber) *Pair {
	left_number.IncreaseDepth()
	right_number.IncreaseDepth()

	return &Pair{
		left_number:  left_number,
		right_number: right_number,
		depth:        INITIAL_DEPTH,
	}
}

func (c *Challenge) LaunchSnailNumberParsing(line string) SnailNumber {
	c.read_index = 0
	c.line = line
	c.current_depth = INITIAL_DEPTH
	return c.ParseSnailNumber()
}

func (c *Challenge) Read() string {
	current_character := string(c.line[c.read_index])
	c.read_index++
	return current_character
}

func (c *Challenge) ParseSnailNumber() SnailNumber {
	var result SnailNumber
	current_character := c.Read()

	if current_character == START_SNAIL_NUMBER_CHARACTER {
		pair_depth := c.current_depth
		c.current_depth++

		left_number := c.ParseSnailNumber()
		current_character = c.Read()
		if current_character != SEPARATION_CHARACTER {
			log.Fatal("should have separation character '", current_character, "' found, read_index=", c.read_index)
		}
		right_number := c.ParseSnailNumber()
		current_character = c.Read()
		if current_character != END_SNAIL_NUMBER_CHARACTER {
			log.Fatal("should have end character '", current_character, "' found, read_index=", c.read_index)
		}
		c.current_depth--

		result = &Pair{
			left_number:  left_number,
			right_number: right_number,
			depth:        pair_depth,
		}
	} else {
		whole_value := ""
		value_character := current_character
		for value_character != SEPARATION_CHARACTER && value_character != END_SNAIL_NUMBER_CHARACTER {
			whole_value += value_character
			value_character = c.Read()
		}

		c.read_index--

		result = &Value{
			value: utils.FatalReadInt(whole_value),
			depth: c.current_depth,
		}
	}

	return result
}
