package challenge

import (
	"fmt"
	"log"
	"math"
)

type Challenge struct {
}

type Node struct {
	x          int
	y          int
	risk       int
	neighbours []*Node
}

func (n Node) String() string {
	return fmt.Sprintf("(%s)->%d", n.Key(), n.risk)
}

func (n Node) Key() string {
	return fmt.Sprintf("%d,%d", n.x, n.y)
}

func (c Challenge) Heuristic(node, end *Node) int {
	return int(math.Abs(float64(end.x-node.x)) + math.Abs(float64(end.y-node.y)))
}

func (c Challenge) Distance(current, neighbour *Node) int {
	return neighbour.risk
}

func (n *Node) IsDifferentFrom(node *Node) bool {
	return n.Key() != node.Key()
}

func (c Challenge) ReconstructPath(cameFrom map[string]*Node, current *Node) []*Node {
	path := []*Node{current}
	ok := true
	for ok {
		current, ok = cameFrom[current.Key()]
		path = append([]*Node{current}, path...)
	}

	return path
}

func (c Challenge) IsIndexValid(index, max_index int) bool {
	return index >= 0 && index < max_index
}

func (c Challenge) FindLowestPath(start, end *Node) []*Node {
	open_set := []*Node{start}
	cameFrom := make(map[string]*Node, 0)

	gScore := make(map[string]int, 0)
	fScore := make(map[string]int, 0)

	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			gScore[fmt.Sprintf("%d,%d", x, y)] = math.MaxInt
			fScore[fmt.Sprintf("%d,%d", x, y)] = math.MaxInt
		}
	}
	gScore[start.Key()] = 0
	fScore[start.Key()] = c.Heuristic(start, end)

	step := 0
	for len(open_set) > 0 {
		log.Println(step)
		var current *Node
		min := math.MaxInt

		for _, node := range open_set {
			if fScore[node.Key()] < min {
				min = fScore[node.Key()]
				current = node
			}
		}

		if !current.IsDifferentFrom(end) {
			return c.ReconstructPath(cameFrom, current)
		}

		new_open_set := make([]*Node, 0)
		for _, node := range open_set {
			if node.IsDifferentFrom(current) {
				new_open_set = append(new_open_set, node)
			}
		}

		open_set = new_open_set

		for _, neighbour := range current.neighbours {
			current_gScore := gScore[current.Key()]
			neighbour_gScore := gScore[neighbour.Key()]
			tentative_gScore := current_gScore + c.Distance(current, neighbour)

			if tentative_gScore < neighbour_gScore {
				cameFrom[neighbour.Key()] = current
				gScore[neighbour.Key()] = tentative_gScore
				fScore[neighbour.Key()] = tentative_gScore + c.Heuristic(neighbour, end)

				is_neighbour_in_open_set := false
				for _, node := range open_set {
					if !node.IsDifferentFrom(neighbour) {
						is_neighbour_in_open_set = true
						break
					}
				}

				if !is_neighbour_in_open_set {
					open_set = append(open_set, neighbour)
				}
			}
		}
		step++
	}

	return []*Node{}
}
