package main

import (
	"2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func find_first_and_last_digit(line string) string {
	digit_words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digit_words_start_index := [10]int{-1}
	digit_words_end_index := [10]int{-1}

	// Find the first and last indexes of each word representation
	for i, word := range digit_words {
		digit_words_start_index[i] = strings.Index(line, word)
		digit_words_end_index[i] = strings.LastIndex(line, word)
	}

	// Find the smallest start index
	min_start_index := math.MaxInt
	start_digit := ""
	for i, index := range digit_words_start_index {
		if index > -1 && index < min_start_index {
			min_start_index = index
			start_digit = strconv.Itoa(i)
		}
	}

	// Find the largest end index
	max_end_index := -1
	end_digit := ""
	for i, index := range digit_words_end_index {
		if index > max_end_index {
			max_end_index = index
			end_digit = strconv.Itoa(i)
		}
	}

	// Find the first and last indexes of each numerical representation
	first_digit_index := strings.IndexAny(line, "1234567890")
	if first_digit_index < min_start_index {
		start_digit = string(line[first_digit_index])
	}

	last_digit_index := strings.LastIndexAny(line, "1234567890")
	if last_digit_index > max_end_index {
		end_digit = string(line[last_digit_index])
	}

	two_digit_number := start_digit + end_digit
	return two_digit_number
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	sum := 0
	for _, line := range lines {
		result, _ := strconv.Atoi(find_first_and_last_digit(line))
		sum += result
	}
	fmt.Println(sum)
}
