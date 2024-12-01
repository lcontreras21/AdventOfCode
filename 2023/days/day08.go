package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Day_8_path struct {
	left  string
	right string
}

type Day_8_start struct {
	key      string
	curr_key string
	steps    int
}

func Day_8_parse_input() (paths map[string]Day_8_path, directions string, ghosts []Day_8_start) {
	file, err := os.Open("2023/inputs/Day_08.txt")
	// file, err := os.Open("2023/inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// Iterate to find all numbers
	i := 0

	paths = make(map[string]Day_8_path)

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if i == 0 {
			directions = txt
		}

		if strings.Contains(txt, " = ") {
			key := txt[:3]
			L := txt[7:10]
			R := txt[12:15]
			paths[key] = Day_8_path{left: L, right: R}

			if key[2:3] == "A" {
				ghosts = append(ghosts, Day_8_start{key: key, curr_key: key, steps: 0})
			}
		}

		i++
	}
	file.Close()
	return
}

func Day_8_Part_1() {
	paths, directions, _ := Day_8_parse_input()

	steps := 0
	dir_i := 0

	done := 0

	curr_key := "AAA"

	for done < 1 {
		if curr_key == "ZZZ" {
			break
		}
		curr_step := directions[dir_i : dir_i+1]
		curr_paths := paths[curr_key]

		if curr_step == "R" {
			curr_key = curr_paths.right
		}
		if curr_step == "L" {
			curr_key = curr_paths.left
		}
		dir_i++
		if dir_i >= len(directions) {
			dir_i = 0
		}
		steps++
	}
	fmt.Println(steps)
}

func Day_8_Part_2() {
	paths, directions, ghosts := Day_8_parse_input()
    fmt.Println(ghosts)

	done := 0
	dir_i := 0
	continue_count := 0

	for done < 1 {
		curr_step := directions[dir_i : dir_i+1]
		continue_count = 0
		for ghost_i, ghost := range ghosts {
			if ghost.curr_key[2:3] == "Z" {
				continue_count++
				continue
			}
			curr_paths := paths[ghost.curr_key]
			if curr_step == "R" {
				ghost.curr_key = curr_paths.right
			}
			if curr_step == "L" {
				ghost.curr_key = curr_paths.left
			}
			ghost.steps++
            ghosts[ghost_i] = ghost
		}
		if continue_count == len(ghosts) {
			break
		}
		dir_i++
		if dir_i >= len(directions) {
			dir_i = 0
		}
	}
    fmt.Println(ghosts)
    var steps []int
    for _, ghost := range ghosts {
        steps = append(steps, ghost.steps)
    }
    lcm := utils.LCMMultiple(steps)
    fmt.Println(lcm)
}
