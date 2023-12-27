package main

import (
	"2023/utils"
	"fmt"
)

func shift_north(board []rune, width int, depth int) {
	for i := 0; i < width; i++ {
		stop_index := i
		for j := 0; j < depth; j++ {
			if board[i+j*width] == '#' {
				stop_index = i + (j+1)*width
			} else if board[i+j*width] == 'O' {
				board[i+j*width] = '.'
				board[stop_index] = 'O'
				stop_index = stop_index + width
			}
		}
	}
}

func shift_west(board []rune, width int, depth int) {
	for i := 0; i < depth; i++ {
		stop_index := i * width
		for j := 0; j < width; j++ {
			if board[i*width+j] == '#' {
				stop_index = i*width + j + 1
			} else if board[i*width+j] == 'O' {
				board[i*width+j] = '.'
				board[stop_index] = 'O'
				stop_index = stop_index + 1
			}
		}
	}
}

func shift_south(board []rune, width int, depth int) {
	for i := 0; i < width; i++ {
		stop_index := i + (depth-1)*width
		for j := depth - 1; j >= 0; j-- {
			if board[i+j*width] == '#' {
				stop_index = i + (j-1)*width
			} else if board[i+j*width] == 'O' {
				board[i+j*width] = '.'
				board[stop_index] = 'O'
				stop_index = stop_index - width
			}
		}
	}
}

func shift_east(board []rune, width int, depth int) {
	for i := 0; i < depth; i++ {
		stop_index := (i+1)*width - 1
		for j := width - 1; j >= 0; j-- {
			if board[i*width+j] == '#' {
				stop_index = i*width + j - 1
			} else if board[i*width+j] == 'O' {
				board[i*width+j] = '.'
				board[stop_index] = 'O'
				stop_index = stop_index - 1
			}
		}
	}
}

func score_board(board []rune, width int, depth int) int {
	score := 0
	for i := 0; i < len(board); i++ {
		if board[i] == 'O' {
			score += depth - i/width
		}
	}
	return score
}
func main() {
	lines := utils.Open_file("puzzle_input.txt")

	board := make([]rune, 0)
	width := 0
	depth := 0
	for _, line := range lines {
		depth++
		width = len(line)
		board = append(board, []rune(line)...)
	}
	cache := make(map[string]int)
	cycle_states := make([]string, 0)
	cycle_len := 0
	cycle_start := -1
	MAX_ITERATIONS := 1000000000
	for i := 1; i <= MAX_ITERATIONS; i++ {
		shift_north(board, width, depth)
		shift_west(board, width, depth)
		shift_south(board, width, depth)
		shift_east(board, width, depth)
		if val, ok := cache[string(board)]; !ok {
			cache[string(board)] = i
		} else {
			cycle_len = i - val
			if cycle_start == -1 {
				cycle_start = val
			}
			cycle_states = append(cycle_states, string(board))
			if len(cycle_states) == cycle_len {
				break
			}
		}
	}

	i := cycle_start
	for ; i+cycle_len <= MAX_ITERATIONS; i += cycle_len {
	}

	score := score_board([]rune(cycle_states[MAX_ITERATIONS-i]), width, depth)
	fmt.Println(score)
}
