package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day_11_parse_input() (image [][]string) {
	// file, err := os.Open("2023/inputs/Day_11.txt")
	file, err := os.Open("2023/inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := fileScanner.Text()

		split := strings.Split(txt, "")
		image = append(image, split)
	}
	file.Close()
	return
}

func insert_row_into_image(image [][]string, indx int) [][]string {
	row := []string{}
	for i := 0; i < len(image[0]); i++ {
		row = append(row, ".")
	}
	if indx == len(image)-1 {
		image = append(image, row)
	} else {
		image = append(image[:indx+1], image[indx:]...)
		image[indx] = row
	}
	return image
}

func insert_col_into_image(image [][]string, indx int) [][]string {
	for row_i, row := range image {
		if indx == len(row)-1 {
			row = append(row, ".")
		} else {
			row = append(row[:indx+1], row[indx:]...)
			row[indx] = "."
		}
		image[row_i] = row
	}
	return image
}

func gravitational_effects(image [][]string) ([][]int) {
	// Iterate through rows
	i := 0
	row_locs := []int{}
	for i < len(image) {
		row := image[i]
		empty_space := 0
		for _, cell := range row {
			if cell == "." {
				empty_space++
			}
		}
		if empty_space == len(row) {
			row_locs = append(row_locs, i)
		}
		i++
	}

	// Iterate through columns
	i = 0
	col_locs := []int{}
	for i < len(image[0]) {
		col := []string{}
		for _, row := range image {
			col = append(col, row[i])
		}

		empty_space := 0
		for _, cell := range col {
			if cell == "." {
				empty_space++
			}
		}
		if empty_space == len(col) {
			col_locs = append(col_locs, i)
		}
		i++
	}
	locs := [][]int{row_locs, col_locs}
	return locs
}

func print_image(image [][]string) {
	for _, row := range image {
		fmt.Println(row)
	}
}

func calculate_distance(coord_a [2]int, coord_b [2]int, space [][]int, grav_expansion_const int) int {
    empty_row_ids := space[0]
    empty_col_ids := space[1]
    
    rows_in_between := 0
    cols_in_between := 0
    for _, row_id := range empty_row_ids {
        if (coord_a[0] < row_id && coord_b[0] > row_id) || (coord_b[0] < row_id && coord_a[0] > row_id) {
            rows_in_between++
        }
    }
    for _, col_id := range empty_col_ids {
        if (coord_a[1] < col_id && coord_b[1] > col_id) || (coord_b[1] < col_id && coord_a[1] > col_id) {
            cols_in_between++
        }
    }

    row_expansion := rows_in_between * grav_expansion_const
    col_expansion := cols_in_between * grav_expansion_const
	height := coord_b[0] - coord_a[0]
	width := coord_b[1] - coord_a[1]
	if height < 0 {
		height = height * -1
	}
	if width < 0 {
		width = width * -1
	}
	distance := width + height + row_expansion + col_expansion - rows_in_between - cols_in_between
	return distance
}

func get_galaxy_coordinates(image [][]string) (coordinates [][2]int) {
	for row_i, row := range image {
		for col_i, col := range row {
			if col == "#" {
				coordinates = append(coordinates, [2]int{row_i, col_i})
			}
		}
	}
	return
}

func do_work(grav_expansion_const int) {
	image := Day_11_parse_input()
	locs := gravitational_effects(image)
	galaxy_coords := get_galaxy_coordinates(image)

	total := 0
	for g_i, galaxy_coord := range galaxy_coords[:len(galaxy_coords)-1] {
		for _, other_coord := range galaxy_coords[g_i+1:] {
			distance := calculate_distance(galaxy_coord, other_coord, locs, grav_expansion_const)
			total = total + distance
		}
	}
	fmt.Println(total)
}

func Day_11_Part_1() {
    do_work(2)
}

func Day_11_Part_2() {
    do_work(1000000)
}
