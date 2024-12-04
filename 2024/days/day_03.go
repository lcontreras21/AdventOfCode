package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func day_03_parse_input(use_test_file bool) ([]string) {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_03.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

    lines := []string{}
	for fileScanner.Scan() {
		txt := fileScanner.Text()
        lines = append(lines, txt)
	}

	file.Close()
    return lines
}

func Day_3_Part_1() {
	// lines := day_03_parse_input(true)
	lines := day_03_parse_input(false)

    total := 0
    for _, line := range(lines) {
        regexp, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
        matches := regexp.FindAllString(line, -1)
        for _, match := range(matches) {
            match = match[4:len(match)-1]
            nums := strings.Split(match, ",")
            left, _ := strconv.Atoi(nums[0])
            right, _ := strconv.Atoi(nums[1])
            total += (left * right)
        }
    }
    fmt.Println("Corrupted Memory Total: " + strconv.Itoa(total))
}

func Day_3_Part_2() {
	// lines := day_03_parse_input(true)
	lines := day_03_parse_input(false)

    total := 0
    enabled := true
    for _, line := range(lines) {
        mul_regexp, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
        mul_matches := mul_regexp.FindAllStringIndex(line, -1)

        do_regexp, _ := regexp.Compile("do\\(\\)")
        do_matches := do_regexp.FindAllStringIndex(line, -1)

        dont_regexp, _ := regexp.Compile("don't\\(\\)")
        dont_matches := dont_regexp.FindAllStringIndex(line, -1)

        for _, match_indx := range(mul_matches) {
            match := line[match_indx[0]: match_indx[1]]
            match = match[4:len(match)-1]

            if len(dont_matches) > 0 {
                if (match_indx[0] > dont_matches[0][0]) {
                    enabled = false
                    dont_matches = append(dont_matches[:0], dont_matches[1:]...)
                }
            }

            if len(do_matches) > 0 {
                if (match_indx[0] > do_matches[0][0]) {
                    enabled = true
                    do_matches = append(do_matches[:0], do_matches[1:]...)
                }
            }

            if enabled {
                nums := strings.Split(match, ",")
                left, _ := strconv.Atoi(nums[0])
                right, _ := strconv.Atoi(nums[1])
                total += (left * right)
            }
        }
    }
    fmt.Println("Corrupted Memory Total: " + strconv.Itoa(total))
}
