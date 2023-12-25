package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type state struct {
	record_index int
	count_index  int
	cur_len      int
}

var memo map[state]int

func compute_possibilities(s state, record []rune, counts []int) int {
	// Look up result if present
	if val, ok := memo[s]; ok {
		return val
	}
	// Reached the end
	if s.record_index == len(record) {
		if s.count_index == len(counts)-1 && s.cur_len == counts[len(counts)-1] {
			return 1
		} else if s.count_index == len(counts) && s.cur_len == 0 {
			return 1
		} else {
			return 0
		}
	}
	ans := 0
	orig_s := s
	s.record_index++
	if record[orig_s.record_index] == '#' {
		s.cur_len++
		ans += compute_possibilities(s, record, counts)
	} else if record[orig_s.record_index] == '.' && s.cur_len == 0 {
		ans += compute_possibilities(s, record, counts)
	} else if record[orig_s.record_index] == '.' && s.count_index < len(counts) && s.cur_len == counts[s.count_index] {
		s.count_index++
		s.cur_len = 0
		ans += compute_possibilities(s, record, counts)
	} else if record[orig_s.record_index] == '?' {
		// Add possibilities if '?' is '#' or '.'
		// '#' case
		ans += compute_possibilities(state{record_index: s.record_index, count_index: s.count_index, cur_len: s.cur_len + 1},
			record, counts)
		// '.' case
		if s.cur_len == 0 {
			ans += compute_possibilities(s, record, counts)
		} else if s.count_index < len(counts) && s.cur_len == counts[s.count_index] {
			ans += compute_possibilities(state{record_index: s.record_index, count_index: s.count_index + 1, cur_len: 0},
				record, counts)
		}
	}

	memo[orig_s] = ans
	return ans
}

func parse_line(line string) int {
	tokens := strings.Fields(line)

	springs_str := strings.Split(tokens[1], ",")
	var springs []int
	for i := 0; i < len(springs_str); i++ {
		val, _ := strconv.Atoi(springs_str[i])
		springs = append(springs, val)
	}
	var record []rune = []rune(tokens[0])

	records5x := make([]rune, 0, len(record)*5+4)
	springs5x := make([]int, 0, len(springs))

	for i := 0; i < 5; i++ {
		springs5x = append(springs5x, springs...)
		if i > 0 {
			records5x = append(records5x, '?')
		}
		records5x = append(records5x, record...)
	}

	return compute_possibilities(state{0, 0, 0}, records5x, springs5x)
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")
	sum := 0
	for _, line := range lines {
		memo = make(map[state]int)
		val := parse_line(line)
		sum += val
	}
	fmt.Println(sum)

}
