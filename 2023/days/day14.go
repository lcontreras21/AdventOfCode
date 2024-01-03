package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day_14_parse_input(use_test_file bool) [][]string {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_14.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	output := [][]string{}
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, "")
		output = append(output, split)
		i++
	}

	file.Close()
	return output
}

func tilt_landscape_up(landscape [][]string) [][]string {
	for col_i := 0; col_i < len(landscape[0]); col_i++ {
		marker := 0
		for row_i := 0; row_i < len(landscape); row_i++ {
			cell := landscape[row_i][col_i]
			if cell == "." {
				continue
			} else if cell == "#" {
				marker = row_i + 1
			} else { // cell is "O"
				replaced := landscape[marker][col_i]
				landscape[marker][col_i] = "O"
				landscape[row_i][col_i] = replaced
				marker++
			}
		}
	}

	return landscape
}

func calculate_load(landscape [][]string) (load int) {
	for row_i, row := range landscape {
		num_Os := 0
		for _, col := range row {
			if col == "O" {
				num_Os++
			}
		}

		weight := len(landscape) - row_i
		load = load + (num_Os * weight)
	}
	return
}

func Day_14_Part_1() {
	// landscape := Day_14_parse_input(true)
	landscape := Day_14_parse_input(false)
	landscape = tilt_landscape_up(landscape)

	load := calculate_load(landscape)
	fmt.Println(load)
}

func cycle(landscape [][]string) [][]string {
	for i := 0; i < 4; i++ {
		landscape = tilt_landscape_up(landscape)
		landscape = utils.RotateMatrixClockwise[string](landscape)
	}

	return landscape
}

func cashify_landscape(landscape [][]string) string {
	key := ""
	for i := 0; i < len(landscape); i++ {
		k := ""
		for _, cell := range landscape[i] {
			k = k + cell
		}
		key = key + k
	}
	return key
}

func calculate_load_at_cycle(landscape [][]string, cycle_at int) int{
	cycles_seen := []string{}
    cycles_seen = append(cycles_seen, cashify_landscape(landscape))
	reverse_lookup := make(map[int]int)
    reverse_lookup[0] = calculate_load(landscape)

	for i := 1; i < cycle_at; i++ {
		landscape = cycle(landscape)
		load := calculate_load(landscape)

		cashed := cashify_landscape(landscape)
		index := utils.FindIndex[string](cycles_seen, cashed)
		if index >= 0 {
			// We found a cycle!
			cycle_length := len(cycles_seen) - index
			final_index := index + (cycle_at-index)%cycle_length

			return reverse_lookup[final_index]
		}
		cycles_seen = append(cycles_seen, cashify_landscape(landscape))
        reverse_lookup[i] = load
	}
	fmt.Println("dont get here")
	return -1
}

func Day_14_Part_2() {
	// landscape := Day_14_parse_input(true)
	landscape := Day_14_parse_input(false)

    load := calculate_load_at_cycle(landscape, 1e9)
	fmt.Println("Load", load)
}
