package challenge

import (
	"fmt"
)

type Pair struct {
	depth        int
	left_number  SnailNumber
	right_number SnailNumber
}

func (p *Pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.left_number.String(), p.right_number.String())
}

func (p *Pair) AddLeftValue(value int) {
	p.left_number.AddLeftValue(value)
}

func (p *Pair) AddRightValue(value int) {
	p.right_number.AddRightValue(value)
}

func (p *Pair) Explode(c *Challenge) (SnailNumber, bool) {
	if p.depth > 4 {
		new_value := &Value{
			value: 0,
			depth: p.depth,
		}
		if c.last_left_value != nil {
			c.last_left_value.value += p.left_number.Magnitude()
		}
		c.right_value = p.right_number.(*Value)
		return new_value, true
	}

	left_number, has_exploded := p.left_number.Explode(c)
	if has_exploded {
		if c.right_value != nil {
			p.right_number.AddLeftValue(c.right_value.value)
			c.right_value = nil
		}
		p.left_number = left_number
		return p, true
	}

	right_number, has_exploded := p.right_number.Explode(c)
	if has_exploded {
		p.right_number = right_number
		return p, true
	}

	return p, false
}

func (p *Pair) Split() (SnailNumber, bool) {
	left_number, has_splitted := p.left_number.Split()
	if has_splitted {
		p.left_number = left_number
		return p, true
	}

	right_number, has_splitted := p.right_number.Split()
	if has_splitted {
		p.right_number = right_number
		return p, true
	}

	return p, false
}

func (p *Pair) Magnitude() int {
	return 3*p.left_number.Magnitude() + 2*p.right_number.Magnitude()
}

func (p *Pair) IncreaseDepth() {
	p.depth++
	p.left_number.IncreaseDepth()
	p.right_number.IncreaseDepth()
}
