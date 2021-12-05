package challenge

import (
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	result := 1
	return result, time.Since(start)
}
