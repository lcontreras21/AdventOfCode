package days

import (
	"AdventOfCode/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Guard struct {
	loc     models.Coord
	dir     models.Bearing
	path    models.Set[models.Vector]
	in_loop bool
}

func (g Guard) Clone() Guard {
    return Guard{
        loc: models.Coord{X: g.loc.X, Y: g.loc.Y},
        dir: g.dir,
        path: g.path.Clone(),
    }
}

func (g Guard) String() string {
	k := ""
	k = k + "Guard " + g.dir.String() + " - " + g.loc.String()
	return k
}

func (g Guard) UniquePositions() models.Set[[2]int] {
	positions := models.Set[[2]int]{}
	for _, loc := range g.path.ToArray() {
		positions.Append([2]int{loc.Loc.X, loc.Loc.Y})
	}
	return positions
}

func day_06_parse_input(use_test_file bool) (models.Matrix[string], Guard) {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_06.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	area_map := models.Matrix[string]{}
	guard := Guard{}
	line := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, "")
		for col_indx, s := range split {
			if s == "^" {
				guard = Guard{loc: models.Coord{X: line, Y: col_indx}, dir: models.North}
				guard.path.Append(models.Vector{
					Loc: models.Coord{X: line, Y: col_indx},
					Dir: models.North,
				})
			}
		}
		area_map.AddRow(area_map.Rows(), split)
		line += 1
	}

	file.Close()
	return area_map, guard
}

func is_obstacle(area_map models.Matrix[string], loc models.Coord) bool {
	value := area_map.Get(loc.X, loc.Y)
	return value == "#"
}

func move_guard(area_map models.Matrix[string], guard Guard) Guard {
	for true {
		// Move guard
		// Check if out of bounds
		move_loc := models.MoveCoord(guard.loc, guard.dir, 1)
		if is_obstacle(area_map, move_loc) {
			guard.dir = guard.dir.TurnClockwiseBy90()
			continue
		} else {
			guard.loc = move_loc
		}

		if !area_map.IsValidCell(move_loc) {
			break
		}

        is_unique := guard.path.Append(models.Vector{
			Loc: models.Coord{X: move_loc.X, Y: move_loc.Y},
			Dir: guard.dir,
		})

        if !is_unique {
            guard.in_loop = true
            break
        }
	}
	return guard
}

func Day_6_Part_1() {
	area_map, guard := day_06_parse_input(true)
	// area_map, guard := day_06_parse_input(false)

	guard = move_guard(area_map, guard)
	unique_positions := guard.UniquePositions()
	fmt.Println("Number of distinct paths visited:", unique_positions.Length())
}

func Day_6_Part_2() {
	// area_map, guard := day_06_parse_input(true)
	area_map, guard := day_06_parse_input(false)

	starting_position := [2]int{guard.loc.X, guard.loc.Y}
    starting_guard := guard.Clone()

	guard = move_guard(area_map, guard)

	// Iterate through guards path
	// For each spot, place an obstacle and rerun the path calculation
	// If that new_path runs into the original path in same direction
	// it is an overlap
	// if it leaves the board, not an overlap

	guard = move_guard(area_map, guard)
	unique_positions := guard.UniquePositions()
    obstructions := 0
	for _, position := range unique_positions.ToArray() {
		if position == starting_position {
			continue
		}

		cloned_area_map := area_map.Clone()
		cloned_area_map.Set(position[0], position[1], "#")

        new_guard := starting_guard.Clone()
        new_guard = move_guard(cloned_area_map, new_guard)
        if new_guard.in_loop {
            obstructions += 1
        }
	}
	fmt.Println("Number of distinct obstructions:", obstructions)
}
