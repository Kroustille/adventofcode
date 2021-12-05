package challenge

import (
	"testing"

	"github.com/Kroustille/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestResolvePart1(t *testing.T) {
	lines := utils.ReadLines("../test_input")

	c := Challenge{}
	result, _ := c.ResolvePart1(lines)

	assert.Equal(t, result, 1)
}
