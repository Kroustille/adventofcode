package challenge

type Challenge struct {
}

func (c Challenge) IsIndexValid(index, max_value int) bool {
	return index >= 0 && index < max_value
}
