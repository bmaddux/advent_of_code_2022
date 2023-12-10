package main

import (
	"2023/utils"
	"fmt"
	"strings"
)

type Node struct {
	left  string
	right string
}

func create_map(lines []string) map[string]Node {
	nodes := make(map[string]Node)

	for _, line := range lines {
		tokens := strings.Fields(line)
		nodes[tokens[0]] = Node{left: string(tokens[2][1 : len(tokens[2])-1]), right: string(tokens[3][:len(tokens[3])-1])}
	}
	return nodes
}

func find_starting_nodes(lines []string) []string {
	ret := make([]string, 0, 100)
	for _, line := range lines {
		line = strings.Fields(line)[0]
		if string(line[len(line)-1]) == "A" {
			ret = append(ret, line)
		}
	}
	return ret
}

func check_completion(node string) bool {
	return node[len(node)-1] == 'Z'
}

func increment_to_step(desired_step int, node_index int, cur_nodes *[]string, num_steps *[]int, nodes *map[string]Node, directions string) bool {
	if node_index == len(*cur_nodes) {
		return true
	}

	for (*num_steps)[node_index] < desired_step {
		switch directions[(*num_steps)[node_index]%len(directions)] {
		case 'L':
			(*cur_nodes)[node_index] = (*nodes)[(*cur_nodes)[node_index]].left
		case 'R':
			(*cur_nodes)[node_index] = (*nodes)[(*cur_nodes)[node_index]].right
		}
		(*num_steps)[node_index]++
	}
	if node_index > 1 {
		fmt.Println("checking node", node_index, "at step", desired_step, (*cur_nodes)[node_index])
	}
	if check_completion((*cur_nodes)[node_index]) {
		return increment_to_step(desired_step, node_index+1, cur_nodes, num_steps, nodes, directions)
	} else {
		return false
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	nodes := create_map(lines[2:])

	directions := lines[0]
	cur_node := "AAA"

	num_step := 0

	for true {
		for i := 0; i < len(directions); i++ {
			switch directions[i] {
			case 'L':
				cur_node = nodes[cur_node].left
			case 'R':
				cur_node = nodes[cur_node].right
			default:
				panic("invalid direction")
			}
			num_step++
			if cur_node == "ZZZ" {
				break
			}
		}
		if cur_node == "ZZZ" {
			break
		}
	}
	fmt.Println(num_step)

	cur_nodes := find_starting_nodes(lines[2:])
	steps_to_first_z := make([]int, len(cur_nodes), 100)

	for i := 0; i < len(cur_nodes); i++ {
		num_completions := 0
		for true {
			switch directions[steps_to_first_z[i]%len(directions)] {
			case 'L':
				cur_nodes[i] = nodes[cur_nodes[i]].left
			case 'R':
				cur_nodes[i] = nodes[cur_nodes[i]].right
			}
			steps_to_first_z[i]++
			if check_completion(cur_nodes[i]) {
				num_completions++
				break
			}
		}
	}

	fmt.Println(LCM(steps_to_first_z[0], steps_to_first_z[1], steps_to_first_z[2:]...))
}
