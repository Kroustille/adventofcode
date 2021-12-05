package challenge

import (
	"strings"

	"github.com/Kroustille/adventofcode/utils"
)

type Challenge struct {
}

type WindVector struct {
	start Coordinates
	end   Coordinates
}

type Coordinates struct {
	x int
	y int
}

func (c Challenge) ParseLine(line string) WindVector {
	splitted_coords := strings.Split(line, " -> ")
	splitted_start_coords := strings.Split(splitted_coords[0], ",")
	splitted_end_coords := strings.Split(splitted_coords[1], ",")

	w := WindVector{
		start: Coordinates{
			x: utils.FatalReadInt(splitted_start_coords[0]),
			y: utils.FatalReadInt(splitted_start_coords[1]),
		},
		end: Coordinates{
			x: utils.FatalReadInt(splitted_end_coords[0]),
			y: utils.FatalReadInt(splitted_end_coords[1]),
		},
	}

	return w
}

func (w WindVector) IsVerticalOrHorizontal() bool {
	return w.start.x == w.end.x || w.start.y == w.end.y
}
