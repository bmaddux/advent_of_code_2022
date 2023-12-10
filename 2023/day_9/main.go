package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func check_all_zero(values []int) bool {
	for _, val := range values {
		if val != 0 {
			return false
		}
	}
	return true
}

func compute_sum(line string) (int, int) {
	tokens := strings.Fields(line)
	numbers := make([][]int, 1)

	for _, token := range tokens {
		val, _ := strconv.Atoi(token)
		numbers[0] = append(numbers[0], val)
	}

	all_zeros := check_all_zero(numbers[0])
	for !all_zeros {
		numbers = append(numbers, make([]int, 0, 10))

		row_before := len(numbers) - 2
		this_row := row_before + 1
		for i := 0; i < len(numbers[row_before])-1; i++ {
			numbers[this_row] = append(numbers[this_row], numbers[row_before][i+1]-numbers[row_before][i])
		}

		all_zeros = check_all_zero(numbers[this_row])
	}

	for i := len(numbers) - 2; i >= 0; i-- {
		numbers[i] = append(numbers[i], numbers[i+1][len(numbers[i+1])-1]+numbers[i][len(numbers[i])-1])
		numbers[i] = append([]int{numbers[i][0] - numbers[i+1][0]}, numbers[i]...)
	}

	return numbers[0][len(numbers[0])-1], numbers[0][0]
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	part_1_sum := 0
	part_2_sum := 0
	for _, line := range lines {
		part_1, part_2 := compute_sum(line)
		part_1_sum += part_1
		part_2_sum += part_2
	}

	fmt.Println(part_1_sum, part_2_sum)
}
