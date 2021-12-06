package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart1(lines []string) (int, time.Duration) {
	start := time.Now()

	winning_numbers, boards := c.Parse(lines)

	winning_number_index := 0
	winning_board_index := -1
	winning_number := 0
	for winning_number_index < len(winning_numbers) && winning_board_index == -1 {
		number := winning_numbers[winning_number_index]

		for board_index, board := range boards {
			boards[board_index] = c.MarkBoard(board, number)
			is_winning := c.IsWinning(boards[board_index])
			if is_winning {
				winning_board_index = board_index
				winning_number = winning_numbers[winning_number_index]
				log.Println("We have a winner ! Board :")
				log.Println(boards[winning_board_index])
				break
			}
		}

		winning_number_index++
	}

	board_result := c.GetBoardResult(boards[winning_board_index])
	log.Println("board_result=", board_result, "winning_number=", winning_number)

	result := board_result * winning_number

	return result, time.Since(start)
}
