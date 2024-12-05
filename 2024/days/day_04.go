package days

import (
	"AdventOfCode/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_04_parse_input(use_test_file bool) models.Matrix[string] {
	var filename string
	if !use_test_file {
		filename = "2024/inputs/Day_04.txt"
	} else {
		filename = "2024/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	puzzle := models.Matrix[string]{}
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		line := strings.Split(txt, "")
		puzzle.AddRow(puzzle.Rows(), line)
	}

	file.Close()
	return puzzle
}

func is_valid_loc(loc models.Coord, height, width int) bool {
	x_pos, y_pos := loc.X, loc.Y
	if x_pos < 0 || x_pos >= height {
		return false
	}
	if y_pos < 0 || y_pos >= width {
		return false
	}
	return true
}

func get_possible_xmas(puzzle models.Matrix[string], loc models.Coord) []string {
	dirs := models.AllDirs()

	texts := []string{}
	for _, dir := range dirs {
		text := []string{}
		for _, offset := range []int{0, 1, 2, 3} {
			new_loc := models.MoveCoord(loc, dir, offset)
			if is_valid_loc(new_loc, puzzle.Rows(), puzzle.Cols()) {
				text = append(text, puzzle.Get(new_loc.X, new_loc.Y))
			}
		}
		if len(text) == 4 {
			formatted := strings.Join(text, "")
			if formatted == "XMAS" || formatted == "SMAX" {
				texts = append(texts, formatted)
			}
		}
	}
	return texts
}

func Day_4_Part_1() {
	// puzzle := day_04_parse_input(true)
	puzzle := day_04_parse_input(false)

	total := 0
	for row_i := 0; row_i < puzzle.Rows(); row_i++ {
		for col_i := 0; col_i < puzzle.Cols(); col_i++ {
			value := puzzle.Get(row_i, col_i)
			if value == "X" {
				texts := get_possible_xmas(puzzle, models.Coord{X: row_i, Y: col_i})
				total += len(texts)
			}
		}
	}
	fmt.Println("Total amount of XMAS appearances: " + strconv.Itoa(total))
}

func get_leg(puzzle models.Matrix[string], loc models.Coord, dir models.Bearing) string {
    text := []string{}
	for _, offset := range []int{-1, 0, 1} {
		new_loc := models.MoveCoord(loc, dir, offset)
		if is_valid_loc(new_loc, puzzle.Rows(), puzzle.Cols()) {
			text = append(text, puzzle.Get(new_loc.X, new_loc.Y))
		}
    }
    return strings.Join(text, "")
}

func is_mas(s string) bool {
    if s == "MAS" || s == "SAM" {
        return true
    }
    return false
}

func get_possible_xdashmas(puzzle models.Matrix[string], loc models.Coord) []string {
    nw_text := get_leg(puzzle, loc, models.NorthWest)
    ne_text := get_leg(puzzle, loc, models.NorthEast)

    if (is_mas(nw_text) && is_mas(ne_text)) {
        return []string{"MAS"}
    }
    return []string{}
}

func Day_4_Part_2() {
	// puzzle := day_04_parse_input(true)
	puzzle := day_04_parse_input(false)

	total := 0
	for row_i := 0; row_i < puzzle.Rows(); row_i++ {
		for col_i := 0; col_i < puzzle.Cols(); col_i++ {
			value := puzzle.Get(row_i, col_i)
			if value == "A" {
				texts := get_possible_xdashmas(puzzle, models.Coord{X: row_i, Y: col_i})
				total += len(texts)
			}
		}
	}
	fmt.Println("Total amount of X-MAS appearances: " + strconv.Itoa(total))
}
