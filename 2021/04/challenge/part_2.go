package challenge

import (
	"log"
	"time"
)

func (c Challenge) ResolvePart2(lines []string) (int, time.Duration) {
	start := time.Now()

	winning_numbers, boards := c.Parse(lines)

	winning_number_index := 0
	winning_number := 0
	last_winning_board := Board{}
	for winning_number_index < len(winning_numbers) && len(boards) > 0 {
		number := winning_numbers[winning_number_index]

		new_boards := make([]Board, 0)
		for board_index, board := range boards {
			marked_board := c.MarkBoard(board, number)
			boards[board_index] = marked_board
			is_winning := c.IsWinning(marked_board)
			if is_winning {
				winning_number = winning_numbers[winning_number_index]
				last_winning_board = marked_board
			} else {
				new_boards = append(new_boards, marked_board)
			}
		}

		log.Println(len(boards), len(new_boards), winning_numbers[winning_number_index])
		boards = new_boards
		winning_number_index++
	}

	log.Println("Last winning board", last_winning_board)
	board_result := c.GetBoardResult(last_winning_board)
	log.Println("board_result=", board_result, "winning_number=", winning_number)

	result := board_result * winning_number

	return result, time.Since(start)
}
