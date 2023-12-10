package main

import (
	"2023/utils"
	"fmt"
)

func create_tunnels(lines []string) ([]rune, int) {
	tunnels := make([]rune, 0)
	row_len := 0
	for _, line := range lines {
		row_len = len(line)
		tunnels = append(tunnels, []rune(line)...)
	}
	return tunnels, row_len
}

func contains(r rune, options []rune) bool {
	for _, val := range options {
		if val == r {
			return true
		}
	}
	return false
}

type distance struct {
	distance int
	visited  bool
}

func is_valid_connection(lhs rune, lhs_index int, rhs rune, rhs_index int) bool {
	switch lhs {
	case 'S':
		switch {
		// Connects on the right
		case rhs_index == lhs_index+1:
			return contains(rhs, []rune{'-', 'J', '7'})
			// Connects on the left
		case rhs_index == lhs_index-1:
			return contains(rhs, []rune{'-', 'F', 'L'})
			// Connects below
		case rhs_index > lhs_index+1:
			return contains(rhs, []rune{'|', 'J', 'L'})
			// Connects above
		case rhs_index < lhs_index-1:
			return contains(rhs, []rune{'|', 'F', '7'})
		default:
			fmt.Println(lhs_index, rhs_index)
			panic("bad connection")
		}
	case '|':
		switch {
		// Connects above
		case rhs_index < lhs_index-1:
			return contains(rhs, []rune{'|', 'F', '7'})
			// Connects below
		case rhs_index > lhs_index+1:
			return contains(rhs, []rune{'|', 'J', 'L'})
		default:
			return false
		}
	case '-':
		switch {
		// Connects left
		case rhs_index == lhs_index-1:
			return contains(rhs, []rune{'-', 'F', 'L'})
			// Connects right
		case rhs_index == lhs_index+1:
			return contains(rhs, []rune{'-', 'J', '7'})
		default:
			return false
		}
	case 'L':
		switch {
		// Connects above
		case rhs_index < lhs_index-1:
			return contains(rhs, []rune{'|', 'F', '7'})
			// Connects right
		case rhs_index == lhs_index+1:
			return contains(rhs, []rune{'-', 'J', '7'})
		default:
			return false
		}
	case 'F':
		switch {
		// Connects right
		case rhs_index == lhs_index+1:
			return contains(rhs, []rune{'-', 'J', '7'})
			// Connects below
		case rhs_index > lhs_index+1:
			return contains(rhs, []rune{'|', 'J', 'L'})
		default:
			return false
		}
	case '7':
		switch {
		// Connects left
		case rhs_index == lhs_index-1:
			return contains(rhs, []rune{'-', 'F', 'L'})
			// Connects below
		case rhs_index > lhs_index+1:
			return contains(rhs, []rune{'|', 'J', 'L'})
		default:
			return false
		}
	case 'J':
		switch {
		// Connects left
		case rhs_index == lhs_index-1:
			return contains(rhs, []rune{'-', 'F', 'L'})
			// Connects above
		case rhs_index < lhs_index-1:
			return contains(rhs, []rune{'|', 'F', '7'})
		default:
			return false
		}
	default:
		panic("bad pipe")
	}
}

func find_next(cur_index int, tunnels []rune, traversal []distance, row_len int) int {
	next_index := -1
	switch {
	// Go up
	case cur_index-row_len > 0 && is_valid_connection(tunnels[cur_index], cur_index, tunnels[cur_index-row_len], cur_index-row_len) && !traversal[cur_index-row_len].visited:
		next_index = cur_index - row_len
		// Go down
	case cur_index+row_len < len(tunnels) && is_valid_connection(tunnels[cur_index], cur_index, tunnels[cur_index+row_len], cur_index+row_len) && !traversal[cur_index+row_len].visited:
		next_index = cur_index + row_len
		// Go left
	case cur_index%row_len-1 >= 0 && is_valid_connection(tunnels[cur_index], cur_index, tunnels[cur_index-1], cur_index-1) && !traversal[cur_index-1].visited:
		next_index = cur_index - 1
		// Go right
	case cur_index%row_len+1 < row_len && is_valid_connection(tunnels[cur_index], cur_index, tunnels[cur_index+1], cur_index+1) && !traversal[cur_index+1].visited:
		next_index = cur_index + 1
	}
	if next_index != -1 {
		traversal[next_index].visited = true
		traversal[next_index].distance = traversal[cur_index].distance + 1
	}
	return next_index
}

func compute_loop(tunnels []rune, row_len int) {
	start_index := 0
	for i := 0; i < len(tunnels); i++ {
		if tunnels[i] == 'S' {
			start_index = i
			break
		}
	}

	traversal := make([]distance, len(tunnels))
	cur_location := start_index

	for true {
		next := find_next(cur_location, tunnels, traversal, row_len)
		if next == -1 {
			break
		} else {
			cur_location = next
		}
	}

	//for i := 0; i < len(traversal); i++ {
	//	fmt.Print(traversal[i])
	//	if (i+1)%row_len == 0 {
	//		fmt.Println()
	//	}
	//}

	fmt.Println(traversal[cur_location].distance/2 + 1)

	//fmt.Println(" ", string(tunnels[start_index-row_len]))
	//fmt.Println(string(tunnels[start_index-1]), string(tunnels[start_index]), string(tunnels[start_index+1]))
	//fmt.Println(" ", string(tunnels[start_index+row_len]))

	// Just to make the loop complete this depends on the input file
	// Should really be set programatically
	tunnels[start_index] = '|'
	traversal[start_index].visited = true

	inside_counter := 0
	for row := 0; row < len(tunnels)/row_len; row++ {
		outside := true
		last_transitional := '|'
		for col := 0; col < row_len; col++ {
			index := row*row_len + col
			if traversal[index].visited {
				inversion := plane_broken(last_transitional, tunnels[index])
				if inversion {
					outside = outside != true
				}
				if tunnels[index] != '-' {
					last_transitional = tunnels[index]
				}
			} else {
				if !outside {
					inside_counter++
				}
			}
		}
	}
	fmt.Println(inside_counter)
}

func plane_broken(last_transitional rune, next_transitional rune) bool {
	switch last_transitional {
	case '|':
		return next_transitional == '|'
	case 'F':
		return next_transitional == 'J'
	case 'L':
		return next_transitional == '7'
	case 'J':
		return next_transitional == '|'
	case '7':
		return next_transitional == '|'
	}
	return false
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	tunnels, row_len := create_tunnels(lines)

	compute_loop(tunnels, row_len)
}
