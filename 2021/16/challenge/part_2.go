package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	c.all_bits = c.BuildBitString(lines[0])
	c.current_bit_index = 0
	c.version_sum = 0
	log.Println(c.all_bits)

	result := c.Decode()
	return result, time.Since(start)
}
