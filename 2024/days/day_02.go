package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_02_parse_input(use_test_file bool) [][]int {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_02.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	reports := [][]int{}
	for fileScanner.Scan() {
		txt := fileScanner.Text()

		split := strings.Split(txt, " ")
		levels := []int{}
		for _, s := range split {
			v, _ := strconv.Atoi(s)
			levels = append(levels, v)
		}
		reports = append(reports, levels)
	}

	file.Close()
	return reports
}

func Day_2_Part_1() {
	// reports := day_02_parse_input(true)
	reports := day_02_parse_input(false)

	safe := 0
	for _, report := range reports {
        is_valid := check_valid(report)
		if is_valid {
			safe += 1
		}
	}
	fmt.Println("Total number of safe reports: " + strconv.Itoa(safe))
}

func Day_2_Part_2() {
	// reports := day_02_parse_input(true)
	reports := day_02_parse_input(false)

	safe := 0
	for _, report := range reports {
        for level_index := range(report) {
            clone := utils.Clone(report)
            clone = append(clone[:level_index], clone[level_index+1:]...)
            is_safe := check_valid(clone) 
            if is_safe {
                safe += 1
                break
            }
        }
	}
	fmt.Println("Total number of safe reports: " + strconv.Itoa(safe))
}

func check_valid(levels []int) (bool) {
    prev := levels[0]
    prev_sign := levels[1] > levels[0]
    valid := true
    for _, level := range levels[1:] {
        curr_diff := prev - level
        curr_sign := level > prev
        abs_diff := utils.Abs(curr_diff)

        if abs_diff < 1 || abs_diff > 3 {
            valid = false
            break
        }
        if prev_sign != (curr_sign) {
            valid = false
            break
        }
        prev = level
        prev_sign = curr_sign
    }
    return valid
}
