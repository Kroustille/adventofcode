package challenge

import "fmt"

type Challenge struct {
}

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) Key() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Challenge) FoldAlongX(all_coordinates map[string]Coordinates, fold_coordinate int) map[string]Coordinates {
	new_coordinates := make(map[string]Coordinates, 0)
	for _, coordinate := range all_coordinates {
		if coordinate.x < fold_coordinate {
			new_coordinates[coordinate.Key()] = coordinate
		} else {
			folded_coordinate := Coordinates{
				x: fold_coordinate - (coordinate.x - fold_coordinate),
				y: coordinate.y,
			}
			new_coordinates[folded_coordinate.Key()] = folded_coordinate
		}
	}

	return new_coordinates
}

func (c Challenge) FoldAlongY(all_coordinates map[string]Coordinates, fold_coordinate int) map[string]Coordinates {
	new_coordinates := make(map[string]Coordinates, 0)
	for _, coordinate := range all_coordinates {
		if coordinate.y < fold_coordinate {
			new_coordinates[coordinate.Key()] = coordinate
		} else {
			folded_coordinate := Coordinates{
				x: coordinate.x,
				y: fold_coordinate - (coordinate.y - fold_coordinate),
			}
			new_coordinates[folded_coordinate.Key()] = folded_coordinate
		}
	}

	return new_coordinates
}
