package challenge

import (
	"log"
	"strings"

	"github.com/Kroustille/adventofcode/utils"
)

type Challenge struct {
}

type Probe struct {
	x          int
	y          int
	x_velocity int
	y_velocity int
}

type Goal struct {
	start_x int
	end_x   int
	start_y int
	end_y   int
}

func (p *Probe) Step() {
	p.x += p.x_velocity
	p.y += p.y_velocity
	if p.x_velocity > 0 {
		p.x_velocity--
	} else if p.x_velocity < 0 {
		p.x_velocity++
	}

	p.y_velocity--
}

func (c Challenge) ParseGoal(line string) Goal {
	splitted_line := strings.Split(line, ", ")
	x_coords := splitted_line[0]
	y_coords := splitted_line[1]
	log.Println(x_coords)
	log.Println(y_coords)
	splitted_x_coords := strings.Split(x_coords[2:], "..")
	splitted_y_coords := strings.Split(y_coords[2:], "..")

	return Goal{
		start_x: utils.FatalReadInt(splitted_x_coords[0]),
		end_x:   utils.FatalReadInt(splitted_x_coords[1]),
		start_y: utils.FatalReadInt(splitted_y_coords[0]),
		end_y:   utils.FatalReadInt(splitted_y_coords[1]),
	}
}

func (g Goal) IsProbeTooFar(p Probe) bool {
	return p.x > g.end_x || p.y < g.start_y
}

func (g Goal) IsProbeInGoal(p Probe) bool {
	return g.start_x <= p.x &&
		p.x <= g.end_x &&
		g.start_y <= p.y &&
		p.y <= g.end_y
}
