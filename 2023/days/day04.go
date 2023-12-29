package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day_4_Part_1() {
	file, err := os.Open("inputs/Day_04.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("[0-9]+") // Part One

	total := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()

		card_points_split := strings.Split(txt, ": ")

		points_split := strings.Split(card_points_split[1], " | ")

		winning_numbers := r.FindAllString(points_split[0], -1)
		given_numbers := r.FindAllString(points_split[1], -1)

		intersection := utils.SetIntersection(winning_numbers, given_numbers)
		power := len(intersection)

		if power > 0 {
			power = power - 1
			total = total + utils.Power(2, power)
		}
	}

	fmt.Println(total)
	file.Close()
}

func Day_4_Part_2() {
	file, err := os.Open("inputs/Day_04.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("[0-9]+") // Part One

	m := make(map[int]int)
	for fileScanner.Scan() {
		txt := fileScanner.Text()

		card_points_split := strings.Split(txt, ": ")

        card_s := r.FindString(card_points_split[0])
        card, _ := strconv.Atoi(card_s)

		points_split := strings.Split(card_points_split[1], " | ")

		winning_numbers := r.FindAllString(points_split[0], -1)
		given_numbers := r.FindAllString(points_split[1], -1)

		intersection := utils.SetIntersection(winning_numbers, given_numbers)
		value := m[card]
		m[card] = value + 1
		for i := range intersection {
			new_card_value := m[card+i+1]
			m[card+i+1] = new_card_value + 1 + value
		}
	}
	total := 0
	for _, count := range m {
		total = total + count
	}

	fmt.Println(total)
	file.Close()
}
