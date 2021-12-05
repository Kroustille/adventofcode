package challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkBoard(t *testing.T) {
	board := Board{
		Cell{
			number: 1,
			marked: false,
		},
		Cell{
			number: 2,
			marked: true,
		},
		Cell{
			number: 5,
			marked: false,
		},
		Cell{
			number: 10,
			marked: false,
		},
		Cell{
			number: 20,
			marked: false,
		},
		Cell{
			number: 23,
			marked: false,
		},
	}

	expected_win_board := Board{
		Cell{
			number: 1,
			marked: false,
		},
		Cell{
			number: 2,
			marked: true,
		},
		Cell{
			number: 5,
			marked: false,
		},
		Cell{
			number: 10,
			marked: true,
		},
		Cell{
			number: 20,
			marked: false,
		},
		Cell{
			number: 23,
			marked: false,
		},
	}

	expected_fail_board := Board{
		Cell{
			number: 1,
			marked: false,
		},
		Cell{
			number: 2,
			marked: true,
		},
		Cell{
			number: 5,
			marked: false,
		},
		Cell{
			number: 10,
			marked: false,
		},
		Cell{
			number: 20,
			marked: false,
		},
		Cell{
			number: 23,
			marked: false,
		},
	}

	c := Challenge{}
	win_board := c.MarkBoard(board, 10)
	fail_board := c.MarkBoard(board, 7)
	assert.Equal(t, expected_win_board, win_board)
	assert.Equal(t, expected_fail_board, fail_board)
	assert.True(t, !board[3].marked)
}

func TestIsWinning(t *testing.T) {
	board := Board{}
	c := Challenge{}
	is_winning := c.IsWinning(board)
	assert.True(t, is_winning)
}

func TestGetBoardResult(t *testing.T) {
	board := Board{
		Cell{
			number: 1,
			marked: false,
		},
		Cell{
			number: 2,
			marked: true,
		},
		Cell{
			number: 5,
			marked: true,
		},
		Cell{
			number: 10,
			marked: false,
		},
		Cell{
			number: 20,
			marked: true,
		},
		Cell{
			number: 23,
			marked: false,
		},
	}

	c := Challenge{}
	result := c.GetBoardResult(board)
	expected_result := 34

	assert.Equal(t, expected_result, result)
}
