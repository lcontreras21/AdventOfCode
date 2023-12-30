package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bearing int64

const (
	North Bearing = iota
	South
	East
	West
)

type Ray struct {
	loc [2]int
	dir Bearing
}

const (
	Height int = 0
	Width  int = 0
)

func Day_16_parse_input(use_test_file bool) (input [][]string) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_16.txt"
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
		split := strings.Split(txt, "")
		input = append(input, split)
	}

	file.Close()
	return
}

func hashify_ray(ray Ray) string {
	k := strconv.Itoa(ray.loc[0]) + "," + strconv.Itoa(ray.loc[1]) + "," + fmt.Sprint(ray.dir)
	return k
}

func is_valid_loc(loc [2]int, height, width int) bool {
	x_pos, y_pos := loc[0], loc[1]
	if x_pos < 0 || x_pos >= height {
		return false
	}
	if y_pos < 0 || y_pos >= height {
		return false
	}
	return true
}

func move_ray(ray Ray) Ray {
	switch dir := ray.dir; dir {
	case North:
		ray.loc = [2]int{ray.loc[0] - 1, ray.loc[1]}
	case South:
		ray.loc = [2]int{ray.loc[0] + 1, ray.loc[1]}
	case East:
		ray.loc = [2]int{ray.loc[0], ray.loc[1] + 1}
	case West:
		ray.loc = [2]int{ray.loc[0], ray.loc[1] - 1}
	}

	return ray
}

func run_light_beam(contraption [][]string, start_loc [2]int, start_dir Bearing) (energized []string) {
	starting_ray := Ray{loc: start_loc, dir: start_dir}

	rotate_dir_slash := map[Bearing]Bearing{North: East, South: West, East: North, West: South}
	rotate_dir_backslash := map[Bearing]Bearing{North: West, South: East, East: South, West: North}

	rays := []Ray{starting_ray}
	for len(rays) > 0 {
		ray := rays[0]
		rays = rays[1:]

		if !is_valid_loc(ray.loc, len(contraption), len(contraption[0])) {
			continue
		}

		hashed_ray := hashify_ray(ray)
		if utils.FindIndex[string](energized, hashed_ray) < 0 {
			energized = append(energized, hashed_ray)
		} else {
			continue
		}

		cell := contraption[ray.loc[0]][ray.loc[1]]
		if cell == "|" && (ray.dir == East || ray.dir == West) {
			north_ray := Ray{loc: ray.loc, dir: North}
			north_ray = move_ray(north_ray)
			south_ray := Ray{loc: ray.loc, dir: South}
			south_ray = move_ray(south_ray)
			rays = append(rays, north_ray, south_ray)
		} else if cell == "-" && (ray.dir == North || ray.dir == South) {
			east_ray := Ray{loc: ray.loc, dir: East}
			east_ray = move_ray(east_ray)
			west_ray := Ray{loc: ray.loc, dir: West}
			west_ray = move_ray(west_ray)
			rays = append(rays, east_ray, west_ray)
		} else {
			if cell == "/" {
				ray.dir = rotate_dir_slash[ray.dir]
			} else if cell == "\\" {
				ray.dir = rotate_dir_backslash[ray.dir]
			}
			ray = move_ray(ray)
			rays = append(rays, ray)
		}
	}

	seen := map[string]bool{}
	for _, vector := range energized {
		s := strings.Split(vector, ",")
		loc := s[:2]
		seen[strings.Join(loc, ",")] = true
	}
	energized = []string{}
	for k := range seen {
		energized = append(energized, k)
	}

	return
}

func print_energized(contraption [][]string, energized []string) {

	matrix := [][]string{}
	for range contraption {
		row := []string{}
		for range contraption[0] {
			row = append(row, " ")
		}
		matrix = append(matrix, row)
	}
	for _, loc := range energized {
		split := strings.Split(loc, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		matrix[x][y] = "#"
	}
	utils.PrintMatrix[string](matrix)
}

func Day_16_Part_1() {
	// contraption := Day_16_parse_input(true)
	contraption := Day_16_parse_input(false)

	energized := run_light_beam(contraption, [2]int{0, 0}, East)
	// print_energized(contraption, energized)
	fmt.Println(len(energized))
}

func Day_16_Part_2() {
	// contraption := Day_16_parse_input(true)
	contraption := Day_16_parse_input(false)

	total := 0
	for col_i := 0; col_i < len(contraption[0]); col_i++ {
		energized := run_light_beam(contraption, [2]int{0, col_i}, South)
		total = utils.Max(total, len(energized))
	}
	for col_i := 0; col_i < len(contraption[0]); col_i++ {
		energized := run_light_beam(contraption, [2]int{len(contraption) - 1, col_i}, North)
		total = utils.Max(total, len(energized))
	}
	for row_i := 0; row_i < len(contraption); row_i++ {
		energized := run_light_beam(contraption, [2]int{row_i, 0}, East)
		total = utils.Max(total, len(energized))
	}
	for row_i := 0; row_i < len(contraption); row_i++ {
		energized := run_light_beam(contraption, [2]int{row_i, len(contraption[0]) - 1}, West)
		total = utils.Max(total, len(energized))
	}
	fmt.Println(total)

}
