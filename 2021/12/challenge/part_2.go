package challenge

import (
	"fmt"
	"strings"
	"time"
)

func (c Challenge) Go2(current_cave *Cave, all_paths []string, previous_path string, already_small_twice bool) []string {
	if current_cave.IsEnd() {
		final_path := fmt.Sprintf("%s,end", previous_path)
		return append(all_paths, final_path)
	}

	if current_cave.IsStart() && strings.Contains(previous_path, "start") {
		return all_paths
	}

	current_cave.times_viewed++

	current_path := ""
	if previous_path == "" {
		current_path = current_cave.code
	} else {
		current_path = fmt.Sprintf("%s,%s", previous_path, current_cave.code)
	}

	for _, adjacent_cave := range current_cave.linked_caves {
		if adjacent_cave.IsSmall() {
			if adjacent_cave.times_viewed == 0 {
				all_paths = c.Go2(adjacent_cave, all_paths, current_path, already_small_twice)
			} else if !already_small_twice {
				all_paths = c.Go2(adjacent_cave, all_paths, current_path, true)
			}
		} else {
			all_paths = c.Go2(adjacent_cave, all_paths, current_path, already_small_twice)
		}
	}

	current_cave.times_viewed--
	return all_paths
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	start_cave := c.BuildCaves(lines)

	all_paths := c.Go2(start_cave, make([]string, 0), "", false)

	result := len(all_paths)
	return result, time.Since(start)
}
