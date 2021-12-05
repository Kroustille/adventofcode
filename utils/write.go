package utils

import (
	"log"
	"time"
)

func PrintResult(part_number int, result interface{}, start time.Time) {
	PrintResultDuration(part_number, result, time.Since(start))
}

func PrintResultDuration(part_number int, result interface{}, time_ellapsed time.Duration) {
	log.Println("Result of part", part_number, ":", result, "in", time_ellapsed)
}
