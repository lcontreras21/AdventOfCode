package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func day_01_parse_input(use_test_file bool) ([]int, []int) {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_01.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
    
    left_list := []int{}
    right_list := []int{}

	for fileScanner.Scan() {
		txt := fileScanner.Text()

        split := strings.Split(txt, "   ")

        left_num, _ := strconv.Atoi(split[0])
        right_num, _ := strconv.Atoi(split[1])

        left_list = append(left_list, left_num)
        right_list = append(right_list, right_num)
	}

	file.Close()
    return left_list, right_list
}

func Day_1_Part_1() {
    // left_list, right_list := day_01_parse_input(true)
    left_list, right_list := day_01_parse_input(false)

    left_list = utils.Sort(left_list)
    right_list = utils.Sort(right_list)

    total := 0
    for indx := range left_list {
        diff := left_list[indx] - right_list[indx]

        diff = utils.Abs(diff)

        total += diff
    }
    fmt.Println("Offset is " + strconv.Itoa(total))
}

func Day_1_Part_2() {
    // left_list, right_list := day_01_parse_input(true)
    left_list, right_list := day_01_parse_input(false)

    counts := utils.Occurrences(right_list)
    total := 0
    for _, num := range(left_list) {
        count, found := counts[num]
        if found {
            total += (num * count)
        }
    }

    fmt.Println("Simmilarity is " + strconv.Itoa(total))
}
