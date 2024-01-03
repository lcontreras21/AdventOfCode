package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Crucible struct {
	loc      [2]int
	dir      Bearing // from day16 TODO move to separate struct file?
	movement int
	cost     int
	path     []string
	ultra    bool
}

func Day_17_parse_input(use_test_file bool) (city [][]int) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_17.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		line := strings.Split(txt, "")
		to_append := []int{}
		for _, letter := range line {
			num, _ := strconv.Atoi(letter)
			to_append = append(to_append, num)
		}
		city = append(city, to_append)
	}

	file.Close()
	return
}

func add_min_queue(crucibles []Crucible, new Crucible) []Crucible {
	i := 0
	for _, crucible := range crucibles {
		if new.cost < crucible.cost {
			break
		}
		i++
	}

	// Insert at index i
	if i == len(crucibles) {
		crucibles = append(crucibles, new)
	} else {
		crucibles = append(crucibles[:i+1], crucibles[i:]...)
		crucibles[i] = new
	}

	return crucibles
}

func hashify_crucible(crucible Crucible) string {
	k := strconv.Itoa(crucible.loc[0]) + "," + strconv.Itoa(crucible.loc[1])
	k = k + "," + fmt.Sprint(crucible.dir)
	k = k + "," + strconv.Itoa(crucible.movement)
	return k
}

func move_crucible(crucible Crucible, dir Bearing) Crucible {
	switch dir {
	case North:
		crucible.loc = [2]int{crucible.loc[0] - 1, crucible.loc[1]}
	case South:
		crucible.loc = [2]int{crucible.loc[0] + 1, crucible.loc[1]}
	case East:
		crucible.loc = [2]int{crucible.loc[0], crucible.loc[1] + 1}
	case West:
		crucible.loc = [2]int{crucible.loc[0], crucible.loc[1] - 1}
	}

	return crucible
}

func get_possible_crucibles(city [][]int, crucible Crucible) (moved []Crucible) {
	opposite_bearing := map[Bearing]Bearing{North: South, South: North, East: West, West: East}
	for _, bearing := range []Bearing{North, South, East, West} {
		if crucible.ultra {
			if crucible.dir == bearing && crucible.movement == 10 {
				// Can't move more than 10 times
				continue
			}
			if crucible.dir != bearing && crucible.movement < 4 {
				// Can't change direction soon enough
				continue
			}
		} else {
			if crucible.dir == bearing && crucible.movement == 3 {
				// Can't move more than 3 times
				continue
			}
		}

		if opposite_bearing[crucible.dir] == bearing {
			// Can't move backwards
			continue
		}

		updated := move_crucible(crucible, bearing)
		if !is_valid_loc(updated.loc, len(city), len(city[0])) {
			// Can't move out of bounds
			continue
		}

		// Update cost
		updated.cost = updated.cost + city[updated.loc[0]][updated.loc[1]]

		// Update movement
		if updated.dir == bearing {
			updated.movement++
		} else {
			updated.movement = 1
		}

		// Update Dir
		updated.dir = bearing

		// Update Path
		hashed := hashify_crucible(crucible)
		updated.path = utils.Clone[string](updated.path)
		updated.path = append(updated.path, hashed)

		moved = append(moved, updated)
	}

	return
}

func print_path(city [][]int, path []string) {
	grid := [][]string{}
	for range city {
		row := []string{}
		for range city[0] {
			row = append(row, " ")
		}
		grid = append(grid, row)
	}

	for _, cell := range path {
		split := strings.Split(cell, ",")
		x, y := split[0], split[1]
		i, _ := strconv.Atoi(x)
		j, _ := strconv.Atoi(y)
		var direction string
		switch split[2] {
		case "0":
			direction = "^"
		case "1":
			direction = "v"
		case "2":
			direction = ">"
		case "3":
			direction = "<"
		}
		grid[i][j] = direction
	}
	utils.PrintMatrix[string](grid)
}

func get_path_for_crucible(city [][]int, go_plus_ultra bool) int {
	seen := []string{}

	possible_east := Crucible{loc: [2]int{0, 0}, dir: East, ultra: go_plus_ultra}
	possible_south := Crucible{loc: [2]int{0, 0}, dir: South, ultra: go_plus_ultra}
	crucibles := []Crucible{possible_east, possible_south}

	x_end, y_end := len(city)-1, len(city[0])-1

	for len(crucibles) > 0 {
		crucible := crucibles[0]
		crucibles = crucibles[1:]

		next := get_possible_crucibles(city, crucible)
		for _, next_crucible := range next {
			hashed_crucible := hashify_crucible(next_crucible)

			if next_crucible.loc[0] == x_end && next_crucible.loc[1] == y_end {
				if (next_crucible.ultra && next_crucible.movement >= 4) || !next_crucible.ultra {
                    next_crucible.path = append(next_crucible.path, hashed_crucible)
                    // print_path(city, next_crucible.path)
                    return next_crucible.cost
				}
			}

			if utils.FindIndex[string](seen, hashed_crucible) < 0 {
				seen = append(seen, hashed_crucible)
				crucibles = add_min_queue(crucibles, next_crucible)
			}
		}
	}

	// Don't get here :(
	return -1
}

// TODO Optimize Part 2s, they currently take 20 - 60  mins to run 

func Day_17_Part_1() {
	// city := Day_17_parse_input(true)
	city := Day_17_parse_input(false) // Takes a hot sec to run btw

	cost := get_path_for_crucible(city, false)
	fmt.Println("cost", cost)
}

func Day_17_Part_2() {
	// city := Day_17_parse_input(true)
	city := Day_17_parse_input(false)

	cost := get_path_for_crucible(city, true)
	fmt.Println("cost", cost)
}
