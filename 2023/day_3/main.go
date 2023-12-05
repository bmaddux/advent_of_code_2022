package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func create_diagram(lines []string) ([]string, []bool, int) {
	diagram := make([]string, 0)
	include_number := make([]bool, 0)

	row_len := 0
	for _, line := range lines {
		row_len = 0
		for _, char := range line {
			diagram = append(diagram, string(char))
			include_number = append(include_number, false)
			row_len++
		}
	}

	return diagram, include_number, row_len
}

func check_left(index int, row_len int, adjacent_points *[]int) {
	if index%row_len-1 >= 0 {
		*adjacent_points = append(*adjacent_points, index-1)
	}
}

func check_right(index int, row_len int, adjacent_points *[]int) {
	if index%row_len+1 < row_len {
		*adjacent_points = append(*adjacent_points, index+1)
	}
}

func get_surrounding(index int, row_len int) []int {
	adjacent_points := make([]int, 0, 8)

	// Check up
	if index-row_len >= 0 {
		check_left(index-row_len, row_len, &adjacent_points)
		check_right(index-row_len, row_len, &adjacent_points)
		adjacent_points = append(adjacent_points, index-row_len)
	}

	// Check current row
	check_left(index, row_len, &adjacent_points)
	check_right(index, row_len, &adjacent_points)

	// Check down
	if index+row_len < row_len*row_len {
		check_left(index+row_len, row_len, &adjacent_points)
		check_right(index+row_len, row_len, &adjacent_points)
		adjacent_points = append(adjacent_points, index+row_len)
	}

	return adjacent_points
}

func find_numbers(diagram *[]string, include_number *[]bool, row_len int) {
	const non_symbols = "0123456789."
	const digits = "0123456789"

	for i, value := range *diagram {
		// This is a symbol
		if !strings.ContainsAny(value, non_symbols) {
			var adjacent_points []int = get_surrounding(i, row_len)
			for _, adjacent_index := range adjacent_points {
				if strings.ContainsAny((*diagram)[adjacent_index], digits) {
					(*include_number)[adjacent_index] = true

					// Traverse left until non-digit or end
					var go_left int = adjacent_index - 1
					for go_left >= 0 && go_left/row_len == adjacent_index/row_len {
						if strings.ContainsAny((*diagram)[go_left], digits) {
							(*include_number)[go_left] = true
						} else {
							break
						}
						go_left -= 1
					}

					// Traverse right until non-digit or end
					var go_right int = adjacent_index + 1
					for go_right < row_len*row_len && go_right/row_len == adjacent_index/row_len {
						if strings.ContainsAny((*diagram)[go_right], digits) {
							(*include_number)[go_right] = true
						} else {
							break
						}
						go_right += 1
					}
				}
			}
		}
	}
}

func check_adjacent_points(digit *Digit, adjacent_points []int, row_len int) bool {
	for _, point := range adjacent_points {
		if (*digit).start <= point &&
			(*digit).end > point {
			return true
		}
	}
	return false
}

func find_gears(diagram *[]string, row_len int, digits *[]Digit) int {
	var sum int = 0
	for i, value := range *diagram {
		if value == "*" {
			var adjacent_points []int = get_surrounding(i, row_len)
			// Check if any of the digits are adjacent
			var adjacent_digits = make([]*Digit, 0, 2)
			for idx, _ := range *digits {
				if check_adjacent_points(&(*digits)[idx], adjacent_points, row_len) {
					adjacent_digits = append(adjacent_digits, &(*digits)[idx])
				}
			}

			if len(adjacent_digits) == 2 {
				sum += adjacent_digits[0].value * adjacent_digits[1].value
			}
		}
	}
	return sum
}

type Digit struct {
	value int
	start int
	end   int
}

func main() {
	diagram, include_number, row_len := create_diagram(utils.Open_file("puzzle_input.txt"))

	find_numbers(&diagram, &include_number, row_len)

	part_nums_digit := make([]Digit, 0, 1000)

	var builder strings.Builder
	for i, valid_digit := range include_number {
		if valid_digit {
			builder.WriteString(diagram[i])
		} else {
			if builder.Len() > 0 {
				digit_str := builder.String()
				val, _ := strconv.Atoi(digit_str)
				digit := Digit{value: val, start: i - len(digit_str), end: i}
				part_nums_digit = append(part_nums_digit, digit)
				builder.Reset()
			}
		}
		if (i+1)%row_len == 0 {
			if builder.Len() > 0 {
				digit_str := builder.String()
				val, _ := strconv.Atoi(digit_str)
				digit := Digit{value: val, start: i + 1 - len(digit_str), end: i + 1}
				part_nums_digit = append(part_nums_digit, digit)
				builder.Reset()
			}
		}
	}

	var sum int = 0
	for _, part_num := range part_nums_digit {
		sum += part_num.value
	}

	// Part 1
	fmt.Println(sum)

	// Part 2
	fmt.Println(find_gears(&diagram, row_len, &part_nums_digit))
}
