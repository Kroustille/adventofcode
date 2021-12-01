package challenge

type Challenge struct {
}

func (c Challenge) CountMeasureIncreases(measures []int) int {
	increase_count := 0
	last_measure := measures[0]
	for _, current_measure := range measures {
		if current_measure > last_measure {
			increase_count++
		}
		last_measure = current_measure
	}

	return increase_count
}
