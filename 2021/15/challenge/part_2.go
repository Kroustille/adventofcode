package challenge

import (
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()
	all_nodes := make([][]*Node, len(lines)*5)

	for i := 0; i < 5; i++ {
		for x, line := range lines {
			all_nodes[x+i*len(line)] = make([]*Node, len(line)*5)
		}
	}

	for x, line := range lines {
		splitted_line := strings.Split(line, "")
		for y, risk := range splitted_line {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					int_risk := utils.FatalReadInt(risk)
					new_risk := (int_risk + i + j) % 10
					if int_risk+i+j > 10 {
						new_risk = (int_risk+i+j)%10 + (int_risk+i+j)/10
					}
					if new_risk == 0 {
						new_risk = 1
					}
					x_coord := x + i*len(lines)
					y_coord := y + j*len(lines[x])
					all_nodes[x_coord][y_coord] = &Node{
						x:    x_coord,
						y:    y_coord,
						risk: new_risk,
					}
				}
			}
		}
	}
	// log.Println(all_nodes)

	max_line_index := len(all_nodes[0])
	max_col_index := len(all_nodes)
	for line_index, node_line := range all_nodes {
		for col_index, node := range node_line {
			if c.IsIndexValid(line_index+1, max_line_index) {
				node.neighbours = append(node.neighbours, all_nodes[line_index+1][col_index])
			}

			if c.IsIndexValid(line_index-1, max_line_index) {
				node.neighbours = append(node.neighbours, all_nodes[line_index-1][col_index])
			}

			if c.IsIndexValid(col_index+1, max_col_index) {
				node.neighbours = append(node.neighbours, all_nodes[line_index][col_index+1])
			}

			if c.IsIndexValid(col_index-1, max_col_index) {
				node.neighbours = append(node.neighbours, all_nodes[line_index][col_index-1])
			}
		}
	}

	start_node := all_nodes[0][0]
	end_node := all_nodes[len(all_nodes)-1][len(all_nodes[0])-1]

	safest_path := c.FindLowestPath(start_node, end_node)
	sum := 0
	for _, node := range safest_path {
		if node != nil {
			sum += node.risk
		}
	}

	result := sum - start_node.risk
	return result, time.Since(start)
}
