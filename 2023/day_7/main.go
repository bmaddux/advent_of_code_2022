package main

import (
	"2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards     string
	hand_type int
	bet       int
}

func make_hand(line string) Hand {
	tokens := strings.Fields(line)
	hand := Hand{cards: tokens[0]}
	// Determine style of hand
	card_count := [13]int{}

	for _, char := range tokens[0] {
		switch char {
		case 'A':
			card_count[12]++
		case 'K':
			card_count[11]++
		case 'Q':
			card_count[10]++
		case 'J':
			card_count[9]++
		case 'T':
			card_count[8]++
		case '9':
			card_count[7]++
		case '8':
			card_count[6]++
		case '7':
			card_count[5]++
		case '6':
			card_count[4]++
		case '5':
			card_count[3]++
		case '4':
			card_count[2]++
		case '3':
			card_count[1]++
		case '2':
			card_count[0]++
		default:
			panic(char)
		}
	}

	// Find the highest count of any one card
	max_count := 0
	for _, count := range card_count {
		if count > max_count {
			max_count = count
		}
	}

	switch {
	case max_count == 5:
		hand.hand_type = 6
	case max_count == 4:
		hand.hand_type = 5
	case max_count == 3:
		// this is at worst a 3 of a kind
		// Check if its also a full house
		hand.hand_type = 3
		for _, count := range card_count {
			if count == 2 {
				hand.hand_type = 4
				break
			}
		}
	case max_count == 2:
		// this is at worst a pair
		// Check if its also a two pair
		hand.hand_type = 1
		pair_counts := 0
		for _, count := range card_count {
			if count == 2 {
				pair_counts++
			}
		}
		if pair_counts == 2 {
			hand.hand_type = 2
		}
	case max_count == 1:
		hand.hand_type = 0
	}

	hand.bet, _ = strconv.Atoi(tokens[1])

	return hand
}

func cmp_hands(a, b string) int {
	runes := [2][]rune{[]rune(a), []rune(b)}
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch {
			case runes[i][j] == 'T':
				runes[i][j] = 'A'
			case runes[i][j] == 'J':
				runes[i][j] = 'B'
			case runes[i][j] == 'Q':
				runes[i][j] = 'C'
			case runes[i][j] == 'K':
				runes[i][j] = 'D'
			case runes[i][j] == 'A':
				runes[i][j] = 'E'
			}
		}
	}
	for j := 0; j < 5; j++ {
		switch {
		case runes[0][j] < runes[1][j]:
			return -1
		case runes[0][j] > runes[1][j]:
			return 1
		}
	}
	return 0
}

func sort_hands(a, b Hand) int {
	switch {
	case a.hand_type < b.hand_type:
		return -1
	case a.hand_type > b.hand_type:
		return 1
	default:
		return cmp_hands(a.cards, b.cards)
	}
}

func compute_winnings(hands []Hand) int {
	winnings := 0

	for i := 0; i < len(hands); i++ {
		winnings += hands[i].bet * (i + 1)
	}
	return winnings
}

func main() {
	lines := utils.Open_file("puzzle_input.txt")

	hands := make([]Hand, 0, 100)
	for _, line := range lines {
		hands = append(hands, make_hand(line))
	}

	slices.SortFunc(hands, sort_hands)
	fmt.Println(hands)

	fmt.Println(compute_winnings(hands))
}
