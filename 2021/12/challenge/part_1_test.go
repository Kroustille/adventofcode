package challenge

import (
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsSmall(t *testing.T) {
	small_cave := Cave{
		code: "a",
	}
	big_cave := Cave{
		code: "BAS",
	}

	assert.True(t, small_cave.IsSmall())
	assert.False(t, big_cave.IsSmall())
}

func TestResolvePart1(t *testing.T) {
	lines := utils.ReadLines("../test_input")

	c := Challenge{}
	result, _ := c.ResolvePart1(lines)

	assert.Equal(t, result, 1)
}
