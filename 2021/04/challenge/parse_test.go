package challenge

import (
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseWinningNumbers(t *testing.T) {
	lines := utils.ReadLines("../parse_input")

	c := Challenge{}
	winning_numbers, _ := c.Parse(lines)

	expected_result := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	assert.Equal(t, winning_numbers, expected_result)
}

func TestParseLine(t *testing.T) {
	line := " 8  2 23  4 24"

	c := Challenge{}
	result := c.ParseLine(line)
	expected_result := []int{8, 2, 23, 4, 24}

	assert.Equal(t, result, expected_result)
}

func TestParseBoards(t *testing.T) {
	lines := utils.ReadLines("../parse_input")

	c := Challenge{}
	_, boards := c.Parse(lines)

	assert.Equal(t, len(boards), 2)
}
