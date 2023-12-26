package main

import (
	"2023/utils"
	"fmt"
)

type Rock struct {
	row   int
	shape string
}

func shift_north(board [][]Rock, num_rows int) int {
	total_score := 0
	for i := 0; i < len(board); i++ {
		stop_index := 0
		for j := 0; j < len(board[i]); j++ {
			if board[i][j].shape == "#" {
				stop_index = board[i][j].row + 1
			} else {
				board[i][j].row = stop_index
				total_score += num_rows - board[i][j].row
				stop_index++
			}
		}
	}
	return total_score
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	board := make([][]Rock, 0)
	width := 0
	depth := 0
	for i, line := range lines {
		depth++
		width = len(line)
		for j := 0; j < width; j++ {
			if len(board) <= j {
				board = append(board, make([]Rock, 0))
			}
			if line[j] != '.' {
				board[j] = append(board[j], Rock{row: i, shape: string(line[j])})
			}
		}
	}
	total_score := shift_north(board, depth)
	fmt.Println("Total score:", total_score)
}
