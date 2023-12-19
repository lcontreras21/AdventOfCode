package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day_7_hand struct {
	cards     string
	bid       int
}

func get_type(hand Day_7_hand, part int) (t int) {
	hand_quant := utils.QuantifyString(hand.cards)

	// 7 situations
	if len(hand_quant) == 1 {
		t = 7 // Five of a kind
	} else if len(hand_quant) == 2 {
		// Two types
		t = 6 // Four of a kind
		for _, count := range hand_quant {
			if count != 4 && count != 1 {
				t = 5 // Full House
				break
			}
		}
	} else if len(hand_quant) == 3 {
		// Two types
		t = 4 // Three of a kind
		for _, count := range hand_quant {
			if count == 2 {
				t = 3 // Two Pair
				break
			}
		}
	} else if len(hand_quant) == 4 {
		t = 2
	} else {
		t = 1
	}
	if part == 2 {
		count, prs := hand_quant['J']
		if prs {
			if t == 6 || t == 5 { // Four of a kind OR Full House
				t = 7 // Upgrade to Five of a kind
			} else if t == 4 { // Three of a kind
				t = 6 // Upgrade to Four of a kind
			} else if t == 3 && count == 2 { // Two Pair and one of the pairs is a J
				t = 6 // Upgrade to Four of a kind
			} else if t == 3 && count == 1 { // Two Pair and the singleton is the J
				t = 5 // Upgrade to Full House
			} else if t == 2 { // One Pair and the pair is a J
				t = 4 // Upgrade to Three of a kind
			} else if t == 1 { // High card
				t = 2 // Upgrade to One Pair
			}
		}
	}
	return
}

func compare_cards(a, b string, part int) (better bool) {
	// Return if b is syntactically bigger than a

	mapping := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	if part == 2 {
		mapping["J"] = 1
	}

	i := 0
	for i < len(a) {
		l := mapping[a[i:i+1]]
		r := mapping[b[i:i+1]]

		if l != r {
			better = l < r
			break
		}
		i++
	}
	return
}

func compare_hands(a, b Day_7_hand, part int) (better bool) {
	// Return true if b is a better hand than a
	type_a := get_type(a, part)
	type_b := get_type(b, part)
	if type_a == type_b {
		better = compare_cards(a.cards, b.cards, part)
	} else {
		better = type_b > type_a
	}
	return
}

func insert_hand(hands []Day_7_hand, hand Day_7_hand, part int) (new_hands []Day_7_hand) {
	// Use binary search to insert hand

	l := 0
	h := len(hands)

	for l < h {
		mid := int(math.Floor(float64((l + h) / 2)))
		if compare_hands(hands[mid], hand, part) {
			l = mid + 1
		} else {
			h = mid
		}
	}
	if len(hands) == l {
		new_hands = append(hands, hand)
	} else {
		new_hands = append(hands[:l+1], hands[l:]...)
		new_hands[l] = hand
	}
	return
}

func parse_input_day_7(part int) (hands []Day_7_hand) {
	file, err := os.Open("inputs/Day_7.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, " ")
		cards := split[0]
		bid := strings.Replace(split[1], " ", "", -1)
		b, _ := strconv.Atoi(bid)

		hand := Day_7_hand{cards: cards, bid: b}
		hands = insert_hand(hands, hand, part)
	}
	file.Close()
	return
}

func Day_7_Part_1() {
	hands := parse_input_day_7(1)

	total := 0
	i := 1
	for _, hand := range hands {
		total = total + (hand.bid * i)
		i++
	}
	fmt.Println(total)
}

func Day_7_Part_2() {
	hands := parse_input_day_7(2)

	total := 0
	i := 1
	for _, hand := range hands {
		total = total + (hand.bid * i)
		i++
	}
	fmt.Println(total)
}
