package utils

import (
	"log"
	"time"
)

func PrintResult(part_number int, result interface{}, start time.Time) {
	log.Println("Result Part", part_number, ":", result, "in", time.Since(start))
}
