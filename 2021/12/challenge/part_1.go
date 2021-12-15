package challenge

import (
	"fmt"
	"log"
	"time"
)

func (c Challenge) Go(current_cave *Cave, all_paths []string, current_path string) []string {
	if current_cave.IsEnd() {
		final_path := fmt.Sprintf("%s,end", current_path)
		log.Println(final_path)
		return append(all_paths, final_path)
	}

	current_cave.already_viewed = true

	if current_path == "" {
		current_path = current_cave.code
	} else {
		current_path = fmt.Sprintf("%s,%s", current_path, current_cave.code)
	}

	for _, adjacent_cave := range current_cave.linked_caves {
		if !adjacent_cave.already_viewed || !adjacent_cave.IsSmall() {
			all_paths = c.Go(adjacent_cave, all_paths, current_path)
		}
	}

	current_cave.already_viewed = false

	return all_paths
}

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	start_cave := c.BuildCaves(lines)

	all_paths := c.Go(start_cave, make([]string, 0), "")
	// _ = c.Go(start_cave, make([]string, 0), "")

	// log.Println(all_paths)
	result := len(all_paths)
	return result, time.Since(start)
}
