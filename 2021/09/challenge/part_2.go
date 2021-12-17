package challenge

import (
	"sort"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

type Point struct {
	height          int
	adjacent_points []*Point
	in_basin        bool
}

func (p *Point) ComputeBasinSize() int {
	if p.in_basin || p.height == 9 {
		return 0
	}
	size := 1
	p.in_basin = true
	for _, adjacent_point := range p.adjacent_points {
		if adjacent_point.height >= p.height {
			size += adjacent_point.ComputeBasinSize()
		}
	}

	return size
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	all_points := make([][]*Point, len(lines))

	for x, line := range lines {
		splitted_line := strings.Split(line, "")
		for _, character := range splitted_line {
			new_point := &Point{
				height:          utils.FatalReadInt(character),
				adjacent_points: []*Point{},
			}
			all_points[x] = append(all_points[x], new_point)
		}
	}

	points := make([]*Point, 0)
	max_x_index := len(all_points)
	max_y_index := len(all_points[0])
	for x, point_line := range all_points {
		for y, point := range point_line {
			if c.IsIndexValid(x-1, max_x_index) {
				point.adjacent_points = append(point.adjacent_points, all_points[x-1][y])
			}

			if c.IsIndexValid(x+1, max_x_index) {
				point.adjacent_points = append(point.adjacent_points, all_points[x+1][y])
			}

			if c.IsIndexValid(y-1, max_y_index) {
				point.adjacent_points = append(point.adjacent_points, all_points[x][y-1])
			}

			if c.IsIndexValid(y+1, max_y_index) {
				point.adjacent_points = append(point.adjacent_points, all_points[x][y+1])
			}

			points = append(points, point)
		}
	}

	lowest_point_heights := make([]int, 0)
	basin_sizes := []int{}
	for _, point := range points {
		is_lowest := true
		for _, adjacent_point := range point.adjacent_points {
			if adjacent_point.height < point.height {
				is_lowest = false
				break
			}
		}

		if is_lowest {
			basin_size := point.ComputeBasinSize()
			basin_sizes = append(basin_sizes, basin_size)

			lowest_point_heights = append(lowest_point_heights, point.height)
		}
	}

	sort.Ints(basin_sizes)
	greatest_index := len(basin_sizes) - 1
	mid_index := len(basin_sizes) - 2
	lowest_index := len(basin_sizes) - 3

	result := basin_sizes[greatest_index] * basin_sizes[mid_index] * basin_sizes[lowest_index]
	return result, time.Since(start)
}
