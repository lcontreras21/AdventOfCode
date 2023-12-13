package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_2_Part_1() {
	file, err := os.Open("inputs/Day_2.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	max_red, max_green, max_blue := 12, 13, 14

	total := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		splits := strings.Split(txt, ": ")

		// Get ID
		id := strings.Split(splits[0], " ")[1]
		fmt.Println(id)

		// Go through dice pulls per game
		pulls := strings.Split(splits[1], "; ")
		passed := 0
		for i := 0; i < len(pulls); i++ {
			pull := strings.Trim(pulls[i], " ")

			m := make(map[string]int)
			m["red"] = 0
			m["green"] = 0
			m["blue"] = 0

			// Go through colors
			dice := strings.Split(pull, ", ")
			for i := 0; i < len(dice); i++ {
				die := strings.Split(dice[i], " ")
				value, _ := strconv.Atoi(die[0])
				m[die[1]] = m[die[1]] + value
			}

			// Validate
			if m["red"] > max_red || m["green"] > max_green || m["blue"] > max_blue {
				break
			}
			passed++
		}
		if passed == len(pulls) {
			to_add, _ := strconv.Atoi(id)
			total = total + to_add
		}
	}
	fmt.Println("total: ", total)
	file.Close()
}

func Day_2_Part_2() {
	file, err := os.Open("inputs/Day_2.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		splits := strings.Split(txt, ": ")

		m := make(map[string]int)
		m["red"] = 0
		m["green"] = 0
		m["blue"] = 0
		// Go through dice pulls per game
		pulls := strings.Split(splits[1], "; ")
		for i := 0; i < len(pulls); i++ {
			pull := strings.Trim(pulls[i], " ")

			// Go through colors
			dice := strings.Split(pull, ", ")
			for i := 0; i < len(dice); i++ {
				die := strings.Split(dice[i], " ")
				value, _ := strconv.Atoi(die[0])
				m[die[1]] = utils.Max(m[die[1]], value)
			}
			if m["red"] == 0 {
				m["red"] = 1
			}
			if m["green"] == 0 {
				m["green"] = 1
			}
			if m["blue"] == 0 {
				m["blue"] = 1
			}
		}
		cubed := m["red"] * m["green"] * m["blue"]
		total = total + cubed
	}
	fmt.Println("total:", total)
	file.Close()
}
