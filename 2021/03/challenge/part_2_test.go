package challenge

import (
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestFilterRatingsWithColumnBit(t *testing.T) {
	lines := []string{
		"00100",
		"11110",
		"10110",
	}

	c := Challenge{}
	result := c.FilterRatingsWithColumnBit(lines, 0, "1")

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "11110", result[0])
	assert.Equal(t, "10110", result[1])
}

func TestResolvePart2(t *testing.T) {
	lines := utils.ReadLines("../test_input")

	c := Challenge{}
	result, _ := c.ResolvePart2(lines)

	assert.Equal(t, 230, result)
}
