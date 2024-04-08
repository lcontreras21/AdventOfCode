package days

import (
	"AdventOfCode/models"
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DigStep struct {
	dir   Bearing
	steps int
	color string
}

func (d *DigStep) FromColor() DigStep {
	str_steps := d.color[1 : len(d.color)-1]
	str_dir := d.color[len(d.color)-1:]
	steps, _ := strconv.ParseInt(str_steps, 16, 64)
	letter_to_dir := map[string]Bearing{"3": North, "1": South, "0": East, "2": West}
    return DigStep{dir: letter_to_dir[str_dir], steps: int(steps), color:d.color}
}

func Day_18_parse_input(use_test_file bool) (steps []DigStep) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_18.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	letter_to_dir := map[string]Bearing{"U": North, "D": South, "R": East, "L": West}

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, " ")

		dir := letter_to_dir[split[0]]
		move, _ := strconv.Atoi(split[1])
		color := split[2][1 : len(split[2])-1]

		steps = append(steps, DigStep{dir: dir, steps: move, color: color})
	}

	file.Close()
	return
}

func move_loc(loc [2]int, dir Bearing, amount int) [2]int {
	switch dir {
	case North:
		loc = [2]int{loc[0] - amount, loc[1]}
	case South:
		loc = [2]int{loc[0] + amount, loc[1]}
	case East:
		loc = [2]int{loc[0], loc[1] + amount}
	case West:
		loc = [2]int{loc[0], loc[1] - amount}
	}
	return loc
}

func follow_digsteps(digsteps []DigStep) (ground models.Matrix[string], points [][2]int) {
	ground.SetNilValue(".")
	ground.AddEmptyColumn(0, 1)
	ground.Set(0, 0, "#")

	curr_loc := [2]int{0, 0}
	for _, digstep := range digsteps {
		for step := 0; step < digstep.steps; step++ {
			curr_loc = move_loc(curr_loc, digstep.dir, 1)

			if !is_valid_loc(curr_loc, ground.Rows(), ground.Cols()) {
				if digstep.dir == North {
					ground.AddEmptyRow(0, 1)
					curr_loc[0] = 0
				} else if digstep.dir == South {
					ground.AddEmptyRow(ground.Rows(), 1)
				} else if digstep.dir == East {
					ground.AddEmptyColumn(ground.Cols(), 1)
				} else { // West
					ground.AddEmptyColumn(0, 1)
					curr_loc[1] = 0
				}
			}
			ground.Set(curr_loc[0], curr_loc[1], "#")
		}
	}
	inverse := map[Bearing]Bearing{North: South, South: North, East: West, West: East}
	for i := len(digsteps) - 1; i >= 0; i-- {
		digstep := digsteps[i]
		dir := inverse[digstep.dir]
		for step := 0; step < digstep.steps; step++ {
			curr_loc = move_loc(curr_loc, dir, 1)
			points = append(points, curr_loc)
		}
	}
	return
}

func depth(points [][2]int) int {
	area := 0
	for i := 0; i < len(points); i++ {
		coord := points[i]
		var coord_comp [2]int
		if i == len(points)-1 {
			coord_comp = points[0]
		} else {
			coord_comp = points[i+1]
		}
		value := (coord[0] * coord_comp[1]) - (coord[1] * coord_comp[0])
		area = area + value
	}
	area = area / 2
	if area < 0 {
		area = area * -1
	}
	interior := area - len(points)/2 + 1
	return interior + len(points)
}

func Day_18_Part_1() {
    // Follows naive approach to create graph showing steps

	// digsteps := Day_18_parse_input(true)
	digsteps := Day_18_parse_input(false)
	ground, points := follow_digsteps(digsteps)
	// _, points := follow_digsteps(digsteps)
	fmt.Println(ground)

	contains := depth(points)
	fmt.Println("Contains", contains)
}

func tie_shoelace(digsteps []DigStep) int {

    loc := [2]int{0, 0}
    area := 0
    perimeter := 0
    for _, digstep := range digsteps {
        digstep = digstep.FromColor()
        new_loc := move_loc(loc, digstep.dir, digstep.steps)
        area = area + (loc[0] * new_loc[1] - new_loc[0] * loc[1])
        perimeter = perimeter + utils.Abs(new_loc[0] - loc[0]) + utils.Abs(new_loc[1] - loc[1])
        loc = new_loc
    }
    return (utils.Abs(area) + perimeter)/ 2 + 1
}

func Day_18_Part_2() {
    // Use the showlace formula to calculate the area instead to massively spead it up

	// digsteps := Day_18_parse_input(true)
	digsteps := Day_18_parse_input(false)

    contains := tie_shoelace(digsteps)
	fmt.Println("Contains", contains)
}
