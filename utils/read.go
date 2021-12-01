package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ConvertLinesToIntArray(lines []string) []int {
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i] = FatalReadInt(line)
	}

	return numbers
}

func FatalReadInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return number
}
