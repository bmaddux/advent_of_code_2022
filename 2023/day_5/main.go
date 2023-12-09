package main

import (
	"2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Map struct {
	dst int
	src int
	len int
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	seeds := make([]int, 0)
	seeds_tokens := strings.Fields(lines[0])
	for i := 1; i < len(seeds_tokens); i++ {
		seed_int, _ := strconv.Atoi(seeds_tokens[i])
		seeds = append(seeds, seed_int)
	}
	conversion_maps := make([][]Map, 0)

	var map_index int
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			map_index++
		} else if strings.Contains(lines[i], "map") {
			conversion_maps = append(conversion_maps, make([]Map, 0))
		} else {
			tokens := strings.Fields(lines[i])
			dst, _ := strconv.Atoi(tokens[0])
			src, _ := strconv.Atoi(tokens[1])
			len, _ := strconv.Atoi(tokens[2])

			m := Map{
				dst: dst,
				src: src,
				len: len,
			}

			conversion_maps[map_index] = append(conversion_maps[map_index], m)
		}
	}
	var lowest_location = math.MaxInt
	for _, seed := range seeds {
		for i := 0; i < len(conversion_maps); i++ {
			// For each of the conversion ranges, find the one that fits
			for j := 0; j < len(conversion_maps[i]); j++ {
				if seed >= conversion_maps[i][j].src && seed < conversion_maps[i][j].src+conversion_maps[i][j].len {
					seed = conversion_maps[i][j].dst + seed - conversion_maps[i][j].src
					break
				}
			}
		}
		if seed < lowest_location {
			lowest_location = seed
		}
	}
	fmt.Println("lowest location: ", lowest_location)

	// Part 2
	lowest_location = math.MaxInt
	for seed_range := 0; seed_range < len(seeds); seed_range += 2 {
		for seed_start := seeds[seed_range]; seed_start < seeds[seed_range]+seeds[seed_range+1]; seed_start++ {
			var seed = seed_start
			for i := 0; i < len(conversion_maps); i++ {
				// For each of the conversion ranges, find the one that fits
				for j := 0; j < len(conversion_maps[i]); j++ {
					if seed >= conversion_maps[i][j].src && seed < conversion_maps[i][j].src+conversion_maps[i][j].len {
						seed = conversion_maps[i][j].dst + seed - conversion_maps[i][j].src
						break
					}
				}
			}
			if seed < lowest_location {
				lowest_location = seed
			}
		}
	}
	fmt.Println(lowest_location)
}
