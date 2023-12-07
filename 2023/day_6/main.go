package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	time := make([]int, 0, 3)
	distance := make([]int, 0, 3)
	num_solves := make([]int, 0, 3)

	times := strings.Fields(lines[3])
	distances := strings.Fields(lines[4])

	fmt.Println(times)
	fmt.Println(distances)

	for i := 1; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		time = append(time, t)
		distance = append(distance, d)
	}

	for i := 0; i < len(time); i++ {
		num := 0
		for hold_time := 1; hold_time < time[i]; hold_time++ {
			distance_traveled := hold_time * (time[i] - hold_time)
			if distance_traveled > distance[i] {
				num++
			}
		}
		num_solves = append(num_solves, num)
	}
	var total int = 1
	for i := 0; i < len(num_solves); i++ {
		total *= num_solves[i]
	}
	fmt.Println(time, distance, num_solves, total)
}
