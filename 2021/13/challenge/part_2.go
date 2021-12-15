package challenge

import (
	"time"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	result := 2
	return result, time.Since(start)
}
