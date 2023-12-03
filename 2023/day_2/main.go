package main

import (
	"2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func check_colors(line string) (bool, int) {
	num_red := 12
	num_green := 13
	num_blue := 14

	game_str := strings.Split(line, ": ")[0]
	game_id, _ := strconv.Atoi(strings.Split(game_str, " ")[1])

	balls := strings.Split(line, ": ")[1]
	ball_draws := strings.Split(balls, "; ")

	for _, draw := range ball_draws {
		valid := true
		draw_colors := strings.Split(draw, ", ")
		for _, count_color := range draw_colors {
			count, _ := strconv.Atoi(strings.Split(count_color, " ")[0])
			color := strings.Split(count_color, " ")[1]

			switch color {
			case "red":
				if count > num_red {
					valid = false
				}
			case "green":
				if count > num_green {
					valid = false
				}
			case "blue":
				if count > num_blue {
					valid = false
				}
			default:
				panic("invalid str")
			}
		}
		if !valid {
			return false, game_id
		}
	}
	return true, game_id
}

func get_power(line string) int {
	balls := strings.Split(line, ": ")[1]
	ball_draws := strings.Split(balls, "; ")

	min_red, min_green, min_blue := 0, 0, 0

	for _, draw := range ball_draws {
		draw_colors := strings.Split(draw, ", ")
		for _, count_color := range draw_colors {
			count, _ := strconv.Atoi(strings.Split(count_color, " ")[0])
			color := strings.Split(count_color, " ")[1]

			switch color {
			case "red":
				if count > min_red {
					min_red = count
				}
			case "green":
				if count > min_green {
					min_green = count
				}
			case "blue":
				if count > min_blue {
					min_blue = count
				}
			default:
				panic("invalid str")
			}
		}
	}
	return min_red * min_green * min_blue
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	id_sum := 0
	power_sum := 0

	for _, line := range lines {
		valid, game_id := check_colors(line)
		if valid {
			id_sum += game_id
		}
		power_sum += get_power(line)
	}

	fmt.Println("Valid game id sum: ", id_sum)
	fmt.Println("Power sum: ", power_sum)
}
