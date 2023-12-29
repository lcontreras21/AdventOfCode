package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	pattern [][]string
}

func print_patterns(patterns []Pattern) {
	for _, pattern := range patterns {
		for _, row := range pattern.pattern {
			fmt.Println(row)
		}
		fmt.Println()
	}
}

func Day_13_parse_input(use_test_file bool) (patterns []Pattern) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_13.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	stored_row := [][]string{}
	i := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if len(txt) == 0 {
			new_pattern := Pattern{pattern: stored_row}
			patterns = append(patterns, new_pattern)
			stored_row = [][]string{}
		} else {
			split := strings.Split(txt, "")
			stored_row = append(stored_row, split)
		}
		i++
	}
	if len(stored_row) > 0 {
		new_pattern := Pattern{pattern: stored_row}
		patterns = append(patterns, new_pattern)
	}

	file.Close()
	return
}

func confirm_mirror(matrix [][]string, index int, part_two bool) bool {
	index++
	x := utils.Min(index, len(matrix)-index)
	first_half := matrix[index-x : index]
	second_half := matrix[index : index+x]
	second_half = utils.Reverse[[]string](second_half)

	if !part_two {
		for i := range first_half {
			if !utils.CompareArrays[string](first_half[i], second_half[i]) {
				return false
			}
		}
		return true
	} else {
        difference := 0
        for i := range first_half {
            d := utils.CompareArraysWithDifference[string](first_half[i], second_half[i])
            difference = difference + d
        }
        return difference == 1
	}
}

func find_mirror(pattern Pattern, is_col, part_two bool) int {
	matrix := pattern.pattern
	if is_col {
		matrix = utils.Transpose[string](matrix)
	}
	for i := 0; i < len(matrix)-1; i++ {
		if confirm_mirror(matrix, i, part_two) {
			return i + 1
		}
	}
	return 0
}

func Day_13_Part_1() {
	patterns := Day_13_parse_input(false)
	part_two := false

	total := 0
	for _, pattern := range patterns {
		row_mirror_index := find_mirror(pattern, false, part_two)
		if row_mirror_index > 0 {
			total = total + 100*row_mirror_index
		} else {
			col_mirror_index := find_mirror(pattern, true, part_two)
			total = total + col_mirror_index
		}
	}
	fmt.Println(total)
}

func Day_13_Part_2() {
	patterns := Day_13_parse_input(false)
	// patterns := Day_13_parse_input(true)

	part_two := true

	total := 0
	for _, pattern := range patterns {
		row_mirror_index := find_mirror(pattern, false, part_two)
		if row_mirror_index > 0 {
			total = total + 100*row_mirror_index
		} else {
			col_mirror_index := find_mirror(pattern, true, part_two)
			total = total + col_mirror_index
		}
	}
	fmt.Println(total)
}
