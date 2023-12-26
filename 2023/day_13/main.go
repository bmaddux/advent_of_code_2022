package main

import (
	"2023/utils"
	"fmt"
)

func check_vertical_mirror(lhs int, mirror []rune, width int) bool {
	rhs := lhs + 1
	num_rows := len(mirror) / width
	for lhs >= 0 && rhs < width {
		for i := 0; i < num_rows; i++ {
			if mirror[i*width+lhs] != mirror[i*width+rhs] {
				return false
			}
		}
		lhs--
		rhs++
	}
	return true
}

func check_horizontal_mirror(top int, mirror []rune, width int) bool {
	bottom := top + 1
	num_rows := len(mirror) / width
	for top >= 0 && bottom < num_rows {
		for i := 0; i < width; i++ {
			if mirror[width*top+i] != mirror[width*bottom+i] {
				return false
			}
		}
		top--
		bottom++
	}
	return true
}

func check_vertical_mirror2(lhs int, mirror []rune, width int) bool {
	rhs := lhs + 1
	num_rows := len(mirror) / width
	num_errors := 0
	for lhs >= 0 && rhs < width {
		for i := 0; i < num_rows; i++ {
			if mirror[i*width+lhs] != mirror[i*width+rhs] {
				num_errors++
			}
		}
		lhs--
		rhs++
	}
	return num_errors == 1
}

func check_horizontal_mirror2(top int, mirror []rune, width int) bool {
	bottom := top + 1
	num_rows := len(mirror) / width
	num_errors := 0
	for top >= 0 && bottom < num_rows {
		for i := 0; i < width; i++ {
			if mirror[width*top+i] != mirror[width*bottom+i] {
				num_errors++
			}
		}
		top--
		bottom++
	}
	return num_errors == 1
}

func process_mirror(mirror []rune, width int, linenum int) int {
	// Check for vertical mirror
	for i := 0; i < width-1; i++ {
		if check_vertical_mirror(i, mirror, width) {
			return i + 1
		}
	}

	// Check for horizontal mirror
	num_rows := len(mirror) / width
	for i := 0; i < num_rows-1; i++ {
		if check_horizontal_mirror(i, mirror, width) {
			return (i + 1) * 100
		}
	}
	fmt.Println("error on", linenum)
	panic("mirror not found")
}

func process_mirror2(mirror []rune, width int, linenum int) int {
	// Check for vertical mirror
	for i := 0; i < width-1; i++ {
		if check_vertical_mirror2(i, mirror, width) {
			return i + 1
		}
	}

	// Check for horizontal mirror
	num_rows := len(mirror) / width
	for i := 0; i < num_rows-1; i++ {
		if check_horizontal_mirror2(i, mirror, width) {
			return (i + 1) * 100
		}
	}
	fmt.Println("error on", linenum)
	panic("mirror not found")
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	mirror := make([]rune, 0, 100)
	last_width := 0
	sum := 0
	sum2 := 0
	for i, line := range lines {
		if len(line) == 0 || i == len(lines)-1 {
			sum += process_mirror(mirror, last_width, i)
			sum2 += process_mirror2(mirror, last_width, i)
			mirror = make([]rune, 0, 100)
			continue
		} else {
			last_width = len(line)
		}
		mirror = append(mirror, []rune(line)...)
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
