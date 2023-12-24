package main

import (
	"2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func validate_record(record []rune, counts []int) bool {
	live_counts := make([]int, 0, len(counts))
	spring_len := 0
	counts_index := 0
	for i := 0; i < len(record); i++ {
		switch record[i] {
		case '#':
			{
				// Adding an extra spring
				if spring_len == 0 && counts_index >= len(counts) {
					return false
				}
				spring_len++
				// Adding to many springs to the current block
				if spring_len > counts[counts_index] {
					return false
				}
			}
		case '?':
			{
				return true
			}
		case '.':
			{
				if spring_len > 0 {
					live_counts = append(live_counts, spring_len)
					counts_index++
					spring_len = 0
				}
			}
		default:
			panic("unknown character")
		}
	}
	if spring_len > 0 {
		live_counts = append(live_counts, spring_len)
	}
	//fmt.Println(counts, live_counts)
	return slices.Equal(counts, live_counts)
}

func compute_possibilities(record []rune, counts []int) int {
	//fmt.Println(string(record))
	if !validate_record(record, counts) {
		return 0
	}
	index := slices.Index(record, '?')
	if index == -1 {
		fmt.Println(string(record), "passes!")
		return 1
	}

	fixed_copy := make([]rune, len(record))
	broken_copy := make([]rune, len(record))

	copy(fixed_copy, record)
	copy(broken_copy, record)

	fixed_copy[index] = '.'
	broken_copy[index] = '#'

	return compute_possibilities(fixed_copy, counts) + compute_possibilities(broken_copy, counts)
}

func parse_line(line string) int {
	tokens := strings.Fields(line)

	springs_str := strings.Split(tokens[1], ",")
	var springs []int
	for i := 0; i < len(springs_str); i++ {
		val, _ := strconv.Atoi(springs_str[i])
		springs = append(springs, val)
	}
	fmt.Println(line)
	var record []rune = []rune(tokens[0])

	return compute_possibilities(record, springs)
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")
	sum := 0
	for _, line := range lines {
		val := parse_line(line)
		//fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)

}
