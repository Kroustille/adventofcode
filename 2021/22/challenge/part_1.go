package challenge

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

type Coordinates struct {
	start int
	end   int
}

type Cube struct {
	is_on bool
	x     Coordinates
	y     Coordinates
	z     Coordinates
}

func (c Challenge) SplitCoordinates(line string) Coordinates {
	coordinates_string := strings.Split(line, "=")[1]
	coordinates := strings.Split(coordinates_string, "..")

	return Coordinates{
		start: utils.FatalReadInt(coordinates[0]),
		end:   utils.FatalReadInt(coordinates[1]),
	}
}

func (c Coordinates) IsValid() bool {
	return c.start >= -50 &&
		c.start <= 50 &&
		c.end >= -50 &&
		c.end <= 50
}

func (c Cube) IsValid() bool {
	return c.x.IsValid() && c.y.IsValid() && c.z.IsValid()
}

func (c Challenge) ParseCube(line string) Cube {
	splitted_line := strings.Split(line, " ")
	state := splitted_line[0]
	all_coordinates := strings.Split(splitted_line[1], ",")
	x := c.SplitCoordinates(all_coordinates[0])
	y := c.SplitCoordinates(all_coordinates[1])
	z := c.SplitCoordinates(all_coordinates[2])

	cube := Cube{
		is_on: state == "on",
		x:     x,
		y:     y,
		z:     z,
	}

	return cube
}

func (c Challenge) BuildKey(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	engine_coordinates := make(map[string]bool, 0)
	for i, line := range lines {
		cube := c.ParseCube(line)
		if !cube.IsValid() {
			continue
		}

		for x := cube.x.start; x <= cube.x.end; x++ {
			for y := cube.y.start; y <= cube.y.end; y++ {
				for z := cube.z.start; z <= cube.z.end; z++ {
					key := c.BuildKey(x, y, z)
					engine_coordinates[key] = cube.is_on
				}
			}
		}

		log.Println(i, "cube done")
	}

	on_coordinates := 0
	for _, is_coordinate_on := range engine_coordinates {
		if is_coordinate_on {
			on_coordinates++
		}
	}

	result := on_coordinates
	return result, time.Since(start)
}
