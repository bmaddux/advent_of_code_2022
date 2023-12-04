package main

import (
	"2023/utils"
	"fmt"
	"strings"
)

func create_diagram(lines []string) {
	const non_symbols = "0123456789."

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

	// Find all symbols
	for i, char := range diagram {
		// This is a symbol, mark all numbers adjacent to it
		if !strings.Contains(non_symbols, char) {

		}
	}

	fmt.Println(row_len)
	fmt.Println(diagram[0])
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	create_diagram(lines)
}
