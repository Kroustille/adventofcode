package challenge

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/Kroustille/adventofcode/utils"
)

func (c Challenge) SortString(str string) string {
	s := []rune(str)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func (c Challenge) ContainsAllSegments(first_number string, second_number string) bool {
	segments := strings.Split(second_number, "")
	for _, segment := range segments {
		if !strings.Contains(first_number, segment) {
			return false
		}
	}

	return true
}

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()
	sum := 0

	for _, line := range lines {
		splitted_line := strings.Split(line, " | ")
		input := splitted_line[0]
		digits := strings.Split(input, " ")

		zero := ""
		one := ""
		four := ""
		six := ""
		seven := ""
		eight := ""
		nine := ""
		possible_zero_six_nine := make([]string, 0)

		for _, digit := range digits {
			switch len(digit) {
			case 2:
				one = digit
				break
			case 4:
				four = digit
				break
			case 3:
				seven = digit
				break
			case 6:
				possible_zero_six_nine = append(possible_zero_six_nine, digit)
				break
			case 7:
				eight = digit
				break
			default:
				break
			}
		}

		one_segments := strings.Split(one, "")
		seven_segments := strings.Split(seven, "")
		eight_segments := strings.Split(eight, "")
		four_segments := strings.Split(four, "")

		possible_zero_nine := make([]string, 0)
		for _, possible_six := range possible_zero_six_nine {
			all_segments := true
			for _, one_segment := range one_segments {
				all_segments = all_segments && strings.Contains(possible_six, one_segment)
			}

			if !all_segments && six == "" {
				six = possible_six
			} else {
				possible_zero_nine = append(possible_zero_nine, possible_six)
			}
		}

		if c.ContainsAllSegments(possible_zero_nine[0], four) {
			nine = possible_zero_nine[0]
			zero = possible_zero_nine[1]
		} else {
			nine = possible_zero_nine[1]
			zero = possible_zero_nine[0]
		}

		encoded_a := ""
		encoded_b := ""
		encoded_c := ""
		encoded_d := ""
		encoded_e := ""
		encoded_f := ""
		encoded_g := ""

		for i := 0; i < len(seven_segments) && encoded_a == ""; i++ {
			current_segment := seven_segments[i]
			is_segment_a := !strings.Contains(one, current_segment)
			if is_segment_a {
				encoded_a = current_segment
			}
		}

		for _, segment := range eight_segments {
			is_c := !strings.Contains(six, segment)
			if is_c {
				encoded_c = segment
			}

			is_e := !strings.Contains(nine, segment)
			if is_e {
				encoded_e = segment
			}

			is_d := !strings.Contains(zero, segment)
			if is_d {
				encoded_d = segment
			}
		}

		for _, segment := range one_segments {
			is_f := encoded_c != segment
			if is_f {
				encoded_f = segment
			}
		}

		for _, segment := range four_segments {
			if segment != encoded_d && segment != encoded_c && segment != encoded_f {
				encoded_b = segment
			}
		}

		for _, segment := range []string{"a", "b", "c", "d", "e", "f", "g"} {
			if !strings.Contains(encoded_a+encoded_b+encoded_c+encoded_d+encoded_e+encoded_f, segment) {
				encoded_g = segment
			}
		}

		possible_values_of_nine := c.SortString(encoded_a + encoded_b + encoded_c + encoded_d + encoded_f + encoded_g)
		possible_values_of_eight := c.SortString(encoded_a + encoded_b + encoded_c + encoded_d + encoded_e + encoded_f + encoded_g)
		possible_values_of_seven := c.SortString(encoded_a + encoded_c + encoded_f)
		possible_values_of_six := c.SortString(encoded_a + encoded_b + encoded_d + encoded_e + encoded_f + encoded_g)
		possible_values_of_five := c.SortString(encoded_a + encoded_b + encoded_d + encoded_f + encoded_g)
		possible_values_of_four := c.SortString(encoded_b + encoded_c + encoded_d + encoded_f)
		possible_values_of_three := c.SortString(encoded_a + encoded_c + encoded_d + encoded_f + encoded_g)
		possible_values_of_two := c.SortString(encoded_d + encoded_a + encoded_c + encoded_e + encoded_g)
		possible_values_of_one := c.SortString(encoded_c + encoded_f)
		possible_values_of_zero := c.SortString(encoded_a + encoded_b + encoded_c + encoded_e + encoded_f + encoded_g)

		possible_values := []string{
			possible_values_of_zero,
			possible_values_of_one,
			possible_values_of_two,
			possible_values_of_three,
			possible_values_of_four,
			possible_values_of_five,
			possible_values_of_six,
			possible_values_of_seven,
			possible_values_of_eight,
			possible_values_of_nine,
		}

		output := splitted_line[1]
		output_digits := strings.Split(output, " ")
		whole_number := ""
		for _, digit := range output_digits {
			value := 0
			for value < len(possible_values) {
				if possible_values[value] == c.SortString(digit) {
					break
				}
				value++
			}
			if value == len(possible_values) {
				log.Fatal("value not found")
			}
			whole_number = fmt.Sprintf("%s%d", whole_number, value)
		}
		sum += utils.FatalReadInt(whole_number)
	}

	result := sum
	return result, time.Since(start)
}
