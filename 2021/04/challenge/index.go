package challenge

import (
	"fmt"
	"strings"

	"github.com/Kroustille/adventofcode/utils"
)

const BOARD_SIZE = 5

type Challenge struct {
}

type Cell struct {
	number int
	marked bool
}

type Board []Cell

func (c Challenge) ParseWinningNumbers(line string) []int {
	winning_number_inputs := strings.Split(line, ",")
	winning_numbers := make([]int, len(winning_number_inputs))
	for i, winning_number_input := range winning_number_inputs {
		winning_numbers[i] = utils.FatalReadInt(winning_number_input)
	}

	return winning_numbers
}

func (c Challenge) ParseLine(line string) []int {
	numbers := make([]int, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		index := i*2 + i

		number_to_parse := line[index : index+2]
		trimmed_number := strings.TrimSpace(number_to_parse)
		number := utils.FatalReadInt(trimmed_number)
		numbers[i] = number
	}

	return numbers
}

func (c Challenge) Parse(lines []string) ([]int, []Board) {
	first_line := lines[0]

	winning_number_inputs := strings.Split(first_line, ",")
	winning_numbers := make([]int, len(winning_number_inputs))
	for i, winning_number_input := range winning_number_inputs {
		winning_numbers[i] = utils.FatalReadInt(winning_number_input)
	}

	boards := make([]Board, 0)
	i := 2
	for i < len(lines) && i+BOARD_SIZE-1 < len(lines) {
		board := make([]Cell, 0)
		for j := i; j < i+BOARD_SIZE; j++ {
			numbers := c.ParseLine(lines[j])
			for _, number := range numbers {
				cell := Cell{
					number: number,
					marked: false,
				}
				board = append(board, cell)
			}
		}

		boards = append(boards, board)
		i += BOARD_SIZE + 1
	}

	return winning_numbers, boards
}

func (c Challenge) MarkBoard(board Board, winning_number int) Board {
	new_board := make([]Cell, len(board))
	for i, cell := range board {
		if cell.number == winning_number {
			cell.marked = true
		}

		new_board[i] = cell
	}
	return new_board
}

func (c Challenge) IsWinning(board Board) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		horizontal_wins := 0
		vertical_wins := 0
		for j := 0; j < BOARD_SIZE; j++ {
			if board[i*BOARD_SIZE+j].marked {
				horizontal_wins++
			}

			if board[i+j*BOARD_SIZE].marked {
				vertical_wins++
			}
		}

		if horizontal_wins == 5 || vertical_wins == 5 {
			return true
		}
	}

	return false
}

func (c Challenge) GetBoardResult(board Board) int {
	sum := 0
	for _, cell := range board {
		if !cell.marked {
			sum += cell.number
		}
	}

	return sum
}

func (b Board) String() string {
	result := ""
	for i, cell := range b {
		if i%BOARD_SIZE == 0 {
			result += "\n"
		}
		if cell.marked {
			result += "X"
		} else {
			result += " "
		}
		if cell.number < 10 {
			result += " "
		}
		result += fmt.Sprintf("%d", cell.number)
		result += " "
	}
	return result
}
