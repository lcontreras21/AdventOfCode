package days

import (
    "AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse_file() (map[int][][]int, map[int][]int) {
	file, err := os.Open("inputs/Day_3.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r_parts, _ := regexp.Compile("[0-9]+")
	r_gears, _ := regexp.Compile("[^0-9.]")

	parts_map := make(map[int][][]int) // Row ID -> [[part_number start_index end_index] ...]
	gears_map := make(map[int][]int)   // Row ID -> [index ...]

	// Iterate to find all numbers
	i := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()

		parts := r_parts.FindAllStringIndex(txt, -1)
		gears := r_gears.FindAllStringIndex(txt, -1)

		for _, part := range parts {
			// part is [start_index length]

			part_number, _ := strconv.Atoi(txt[part[0]:part[1]])
			part = append([]int{part_number}, part...)
			part[2] = part[2] - 1

			p_list, prs := parts_map[i]
			var entry [][]int
			if prs {
				entry = append(p_list, part)
			} else {
				entry = [][]int{part}
			}
			parts_map[i] = entry
		}
		for _, gear := range gears {
			var entry []int
			if len(gears_map[i]) == 0 {
				entry = []int{gear[0]}
			} else {
				entry = append(gears_map[i], gear[0])
			}
			gears_map[i] = entry
		}
		i++
	}
	file.Close()
	return parts_map, gears_map
}

func Day_3_Part_1() {
	parts_map, gears_map := parse_file()
	total := 0

	// Process Gears
	for row, gears := range gears_map {
		// fmt.Println("row", row)
		for _, gear := range gears {
			// fmt.Println("\tgear", gear)
			// Go through previous row
			parts, _ := parts_map[row-1]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tprevious part", part[0])
					total += part[0]
				}

			}
			// Go through current row
			parts, _ = parts_map[row]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tcurrent part", part[0])
					total += part[0]
				}
			}

			// Go through next row
			parts, _ = parts_map[row+1]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tnext part", part[0])
					// fmt.Println("\t\tgear_range", gear_range, "part_range", part[1:])
					total += part[0]
				}
			}
		}
	}
	fmt.Println(total)
}

func Day_3_Part_2() {
	parts_map, gears_map := parse_file()

	total := 0
	// Process Gears
	for row, gears := range gears_map {
		// fmt.Println("row", row)
		for _, gear := range gears {
			num_adjacent := 0
			var adjacent []int
			// fmt.Println("\tgear", gear)

			// Go through previous row
			parts, _ := parts_map[row-1]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tprevious part", part[0])
					num_adjacent++
					adjacent = append(adjacent, part[0])
				}

			}
			// Go through current row
			parts, _ = parts_map[row]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tcurrent part", part[0])
					num_adjacent++
					adjacent = append(adjacent, part[0])
				}
			}

			// Go through next row
			parts, _ = parts_map[row+1]
			for _, part := range parts {
				gear_range := []int{gear - 1, gear + 1}
				if utils.RangeOverlap(gear_range, part[1:]) {
					// fmt.Println("\t\tnext part", part[0])
					// fmt.Println("\t\tgear_range", gear_range, "part_range", part[1:])
					num_adjacent++
					adjacent = append(adjacent, part[0])
				}
			}
			if num_adjacent == 2 {
				total = total + (adjacent[0] * adjacent[1])
			}
		}
	}
	fmt.Println(total)
}
