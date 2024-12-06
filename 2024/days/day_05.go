package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_05_parse_input(use_test_file bool) (map[int][]int, [][]int) {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_05.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	rules := map[int][]int{}
	updates := [][]int{}
	do_updates := false
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt == "" {
			do_updates = true
			continue
		}

		if !do_updates {
			// Handle rules
			split := strings.Split(txt, "|")
			page_1, _ := strconv.Atoi(split[0])
			page_2, _ := strconv.Atoi(split[1])

			curr, exists := rules[page_1]
			if exists {
				rules[page_1] = append(curr, page_2)
			} else {
				rules[page_1] = []int{page_2}
			}
		} else {
			// Handle updates
			split := strings.Split(txt, ",")
			update := []int{}
			for _, s := range split {
				num, _ := strconv.Atoi(s)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	file.Close()
	return rules, updates
}

func valid_page(to_check, rules []int) int {
	index := -1
	for _, prev_page := range to_check {
		// Check that page is following all of its rules
		rule_index := utils.FindIndex(rules, prev_page)
		if rule_index >= 0 {
			index = rule_index
			break
		}
	}
	return index
}

func get_valid_updates(rules map[int][]int, updates [][]int) ([][]int, [][]int) {
	correctly_ordered := [][]int{}
	incorrectly_ordered := [][]int{}
	for _, update := range updates {
		follows_rules := true
		for curr_indx, page := range update {
			page_rules, exists := rules[page]
			if !exists {
				continue
			}
			found_index := valid_page(update[:curr_indx], page_rules)
			if found_index >= 0 {
				follows_rules = false
				break
			}

			if !follows_rules {
				break
			}
		}
		if follows_rules {
			correctly_ordered = append(correctly_ordered, update)
		} else {
			incorrectly_ordered = append(incorrectly_ordered, update)
		}
	}
	return correctly_ordered, incorrectly_ordered
}

func get_total(updates [][]int) int {
	total := 0
	for _, update := range updates {
		middle_index := (len(update) - 1) / 2
		middle := update[middle_index]
		total += middle
	}
	return total
}

func Day_5_Part_1() {
	// rules, updates := day_05_parse_input(true)
	rules, updates := day_05_parse_input(false)

	correctly_ordered, _ := get_valid_updates(rules, updates)

	total := get_total(correctly_ordered)
	fmt.Println("Number of correctly ordered updates is:", total)
}

func Day_5_Part_2() {
	// rules, updates := day_05_parse_input(true)
	rules, updates := day_05_parse_input(false)

	_, invalid_updates := get_valid_updates(rules, updates)

	valid_updates := [][]int{}
	for update_index := range invalid_updates {
		update := invalid_updates[update_index]
		incorrectly_ordered := 0
		index := 0
		for incorrectly_ordered != 1 {
			if index >= len(update) {
				incorrectly_ordered = 1
				break
			}
			page := update[index]
			page_rules, exists := rules[page]
			if !exists {
				index += 1
				continue
			}
			rule_index := valid_page(update[:index], page_rules)
			if rule_index >= 0 {
				swap_index := utils.FindIndex(update, page_rules[rule_index])
				// Swap and reset
				update[index] = update[swap_index]
				update[swap_index] = page
				index = 0
				continue
			}
			index += 1
		}
		valid_updates = append(valid_updates, update)
	}
	total := get_total(valid_updates)
	fmt.Println("Number of correctly ordered updates is:", total)
}
