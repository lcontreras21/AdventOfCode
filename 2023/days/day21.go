package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
	loc   [2]int
	steps int
}

func Day_21_parse_input(use_test_file bool) (output [][]string, start [2]int) {
	var filename string
	if !use_test_file {
		filename = "2023/inputs/Day_21.txt"
	} else {
		filename = "2023/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		j := strings.Index(txt, "S")
		if j >= 0 {
			start = [2]int{i, j}
		}
		split := strings.Split(txt, "")
		output = append(output, split)
		i++
	}

	file.Close()
	return
}

func get_neighbors(garden [][]string, loc [2]int, is_part_two bool) (neighbors [][2]int) {
	for _, dir := range []Bearing{North, South, West, East} {
		new_loc := move_loc(loc, dir, 1)
		if is_valid_loc(new_loc, len(garden), len(garden[0])) || is_part_two {
			neighbors = append(neighbors, new_loc)
		}
	}
	return
}

func correctify_loc(loc [2]int, height, width int) [2]int {
	if loc[0] < 0 || loc[0] >= height {
		loc[0] = utils.ModLikePython(loc[0], height)
	}
	if loc[1] < 0 || loc[1] >= width {
		loc[1] = utils.ModLikePython(loc[1], width)
	}
	return loc
}

func take_steps(garden [][]string, start [2]int, limit int, is_part_two bool) int {
	set := map[[2]int]bool{start: true}

	for i := 0; i < limit; i++ {
		new_set := map[[2]int]bool{}
		for pos := range set {
			neighbors := get_neighbors(garden, pos, is_part_two)
			for i := 0; i < len(neighbors); i++ {
				neighbor := neighbors[i]
				cloned := utils.Clone[int](neighbor[:])
				actual_loc := [2]int{cloned[0], cloned[1]}
				if is_part_two {
					actual_loc = correctify_loc(actual_loc, len(garden), len(garden[0]))
				}

				if garden[actual_loc[0]][actual_loc[1]] == "." || garden[actual_loc[0]][actual_loc[1]] == "S" {
					new_set[neighbor] = true
				}
			}
		}
		set = new_set
	}

	return len(set)
}

func Day_21_Part_1() {
	// garden, start := Day_21_parse_input(true)
	garden, start := Day_21_parse_input(false)

	limit := 64
	total := take_steps(garden, start, limit, false)
	fmt.Println(total)
}

func Day_21_Part_2() {
	// garden, start := Day_21_parse_input(true)
	garden, start := Day_21_parse_input(false)

	// This works up to a certain point, takes too long after that
	// limit := 5000
	// total := take_steps(garden, start, limit, true)
	// fmt.Println(total)

	// Apparently:
	// The input is a square with an odd numbered size
	// The starting location is in the perfect middle of the square
	// All tiles on the starting row are gardens
	// All tiles on the starting column are gardens

	// https://nickymeuleman.netlify.app/garden/aoc2023-day21
	// https://www.reddit.com/r/adventofcode/comments/18nevo3/2023_day_21_solutions/kef317h/

	// So everytime we get back to the center by moving len(garden) steps down/left/up/right
	// we are back to square one with the amount of new plots we can see
	// that's why it increases exponentially

	// 612941134797232

	// f(t) = at^2 + bt + c
	// where t = to_edge + x * size
	// If we get values at t = 0, 1, 2, we can use Gaussian Elimination to find a, b, c
	// and then plug in (limit - 65) / size for t to get the final result

	// Guassian Elmination
	// f(2) = 4a + 2b + c = R2
	// f(1) = a + b + c   = R1
	// f(0) = c           = R0
	// =>
	// a = 0.5 R2 - R1 + 0.5 R0
	// b = -0.5 R2 + 2 R1 - 1.5 R0
	// c = R0
	limit := 26501365
	value := (limit - 65) / len(garden) // convert limit to what it would be in t-space
	size := len(garden)
	to_edge := int(math.Ceil(float64(size / 2)))
	fn_results := []float64{}

	for _, x := range []int{0, 1, 2} {
		t := to_edge + x*size
        fn_results = append(fn_results, float64(take_steps(garden, start, t, true)))
	}
	R2, R1, R0 := fn_results[2], fn_results[1], fn_results[0]
	a := int(0.5*R2 - R1 + 0.5*R0)
	b := int(-0.5*R2 + 2*R1 - 1.5*R0)
	c := int(R0)

	total := a*(utils.Power(value, 2)) + b*value + c
    fmt.Println(total)
}
