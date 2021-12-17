package challenge

import (
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestResolvePart2(t *testing.T) {
	lines := utils.ReadLines("../test_input")

	c := Challenge{}
	result, _ := c.ResolvePart2(lines)

	assert.Equal(t, result, 2)
}
