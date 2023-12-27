package main

import (
	"2023/utils"
	"fmt"
	"strings"
)

func hash(input string) int {
	val := 0
	for i := 0; i < len(input); i++ {
		val += int(input[i])
		val *= 17
		val %= 256
	}
	return val
}

type Lense struct {
	label string
	value int
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")
	tokens := strings.Split(lines[0], ",")
	boxes := make([][]Lense, 256)
	total := 0

	for _, token := range tokens {
		part_1 := hash(token)
		label_end := -1
		for i := 0; i < len(token); i++ {
			if token[i] == '-' || token[i] == '=' {
				label_end = i
			}
		}
		label := token[:label_end]
		h := hash(label)
		if token[len(token)-1] == '-' {
			remove_index := -1
			for i := 0; i < len(boxes[h]); i++ {
				if boxes[h][i].label == label {
					remove_index = i
				}
			}
			if remove_index >= 0 {
				boxes[h] = append(boxes[h][:remove_index], boxes[h][remove_index+1:]...)
			}
		} else {
			swap_index := -1
			for i := 0; i < len(boxes[h]); i++ {
				if boxes[h][i].label == label {
					swap_index = i
				}
			}
			if swap_index >= 0 {
				boxes[h][swap_index].value = int(token[len(token)-1] - 48)
			} else {
				boxes[h] = append(boxes[h], Lense{label: label, value: int(token[len(token)-1] - 48)})
			}
		}
		total += part_1
	}
	fmt.Println(total)

	part_2 := 0
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes[i]); j++ {
			part_2 += (i + 1) * (j + 1) * boxes[i][j].value
		}
	}
	fmt.Println(part_2)
}
