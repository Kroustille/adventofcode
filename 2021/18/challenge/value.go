package challenge

import (
	"fmt"
	"math"
)

type Value struct {
	value int
	depth int
}

func (v *Value) String() string {
	return fmt.Sprintf("%d", v.value)
}

func (v *Value) Split() (SnailNumber, bool) {
	if v.value >= 10 {
		splitted_value := float64(v.value) / 2
		left_number_value := math.Floor(splitted_value)
		right_number_value := math.Ceil(splitted_value)

		left_value := &Value{
			value: int(left_number_value),
			depth: v.depth + 1,
		}
		right_value := &Value{
			value: int(right_number_value),
			depth: v.depth + 1,
		}
		new_pair := &Pair{
			left_number:  left_value,
			right_number: right_value,
			depth:        v.depth,
		}
		return new_pair, true
	}

	return v, false
}

func (v *Value) AddRightValue(value int) {
	v.value += value
}

func (v *Value) AddLeftValue(value int) {
	v.value += value
}

func (v *Value) Magnitude() int {
	return v.value
}

func (v *Value) IncreaseDepth() {
	v.depth++
}

func (v *Value) Explode(c *Challenge) (SnailNumber, bool) {
	c.last_left_value = v
	return v, false
}
