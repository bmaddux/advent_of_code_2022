package main

import (
	"2023/utils"
	"fmt"
	"strings"
)

// Returns the number of points this card is worth
func contains(list []string, val string) bool {
	for i := 0; i < len(list); i += 1 {
		if list[i] == val {
			return true
		}
	}
	return false
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func parse_card(line string) int {
	numbers := strings.Split(line, ":")[1]

	winning_numbers_str := strings.TrimSpace(strings.Split(numbers, "|")[0])
	our_numbers_str := strings.TrimSpace(strings.Split(numbers, "|")[1])

	winning_numbers := strings.Split(winning_numbers_str, " ")
	for i := 0; i < len(winning_numbers); i += 1 {
		winning_numbers[i] = strings.TrimSpace(winning_numbers[i])
	}
	our_numbers := strings.Split(our_numbers_str, " ")
	for i := 0; i < len(our_numbers); i += 1 {
		our_numbers[i] = strings.TrimSpace(our_numbers[i])
	}

	pow := 0
	for i := 0; i < len(winning_numbers); i += 1 {
		if len(winning_numbers[i]) == 0 {
			continue
		}
		if contains(our_numbers, winning_numbers[i]) {
			pow += 1
		}
	}
	if pow > 0 {
		return IntPow(2, pow-1)
	}
	return 0
}

func parse_cards2(line string) int {
	numbers := strings.Split(line, ":")[1]

	winning_numbers_str := strings.TrimSpace(strings.Split(numbers, "|")[0])
	our_numbers_str := strings.TrimSpace(strings.Split(numbers, "|")[1])

	winning_numbers := strings.Split(winning_numbers_str, " ")
	for i := 0; i < len(winning_numbers); i += 1 {
		winning_numbers[i] = strings.TrimSpace(winning_numbers[i])
	}
	our_numbers := strings.Split(our_numbers_str, " ")
	for i := 0; i < len(our_numbers); i += 1 {
		our_numbers[i] = strings.TrimSpace(our_numbers[i])
	}
	matches := 0
	for i := 0; i < len(winning_numbers); i += 1 {
		if len(winning_numbers[i]) == 0 {
			continue
		}
		if contains(our_numbers, winning_numbers[i]) {
			matches += 1
		}
	}
	return matches
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	var sum int = 0
	for _, line := range lines {
		val := parse_card(line)
		sum += val
	}
	fmt.Println("Part 1: ", sum)

	var card_counts = make([]int, len(lines))
	for i := 0; i < len(card_counts); i++ {
		card_counts[i] = 1
	}

	for i := 0; i < len(card_counts); i++ {
		var num_won = parse_cards2(lines[i])
		for j := 1; j <= num_won; j++ {
			card_counts[i+j] += card_counts[i]
		}
	}

	sum = 0
	for i := 0; i < len(card_counts); i++ {
		sum += card_counts[i]
	}

	fmt.Println("Part 2: ", sum)
}
