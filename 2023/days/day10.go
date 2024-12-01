package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tile int64

// Enums https://www.sohamkamani.com/golang/enums/
const (
	NorthSouth Tile = iota
	EastWest
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	Start
)

type Node struct {
	weight int
	tile   Tile
}

func (node Node) String() string {
	switch node.tile {
	case NorthSouth:
		return "║"
	case EastWest:
		return "═"
	case NorthEast:
		return "╚"
	case NorthWest:
		return "╝"
	case SouthEast:
		return "╔"
	case SouthWest:
		return "╗"
	case Ground:
		return "."
	case Start:
		return "S"
	}
	return ""
}

func (node *Node) from(tile string) Tile {
	var tile_type Tile
	switch tile {
	case "|":
		tile_type = NorthSouth
	case "-":
		tile_type = EastWest
	case "L":
		tile_type = NorthEast
	case "J":
		tile_type = NorthWest
	case "7":
		tile_type = SouthWest
	case "F":
		tile_type = SouthEast
	case ".":
		tile_type = Ground
	case "S":
		tile_type = Start
	}
	return tile_type
}

func contains_tile(t Tile, tiles []Tile) bool {
	for _, tile := range tiles {
		if t == tile {
			return true
		}
	}
	return false
}

func Day_10_parse_input() (input [][]Node, start [2]int) {
	// file, err := os.Open("2023/inputs/Day_10.txt")
	file, err := os.Open("2023/inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	row := 0
	col := -1
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		chars := strings.Split(txt, "")

		if strings.Contains(txt, "S") {
			col = strings.Index(txt, "S")
			start = [2]int{row, col}
		}

		nodes := []Node{}
		for _, char := range chars {
			new_node := Node{}
			new_node.tile = new_node.from(char)
			nodes = append(nodes, new_node)
		}
		input = append(input, nodes)
		row++
	}
	file.Close()
	return
}

func check_north(maze [][]Node, row, col int, neighbors [][2]int) [][2]int {
	if row > 0 {
		node := maze[row-1][col]
		valid_tile_types := []Tile{NorthSouth, SouthWest, SouthEast, Start}
		if contains_tile(node.tile, valid_tile_types) {
			neighbors = append(neighbors, [2]int{row - 1, col})
		}
	}
	return neighbors
}

func check_south(maze [][]Node, row, col int, neighbors [][2]int) [][2]int {
	if row < len(maze) {
		node := maze[row+1][col]
		valid_tile_types := []Tile{NorthSouth, NorthWest, NorthEast, Start}
		if contains_tile(node.tile, valid_tile_types) {
			neighbors = append(neighbors, [2]int{row + 1, col})
		}
	}
	return neighbors
}

func check_east(maze [][]Node, row, col int, neighbors [][2]int) [][2]int {
	if col < len(maze[row]) {
		node := maze[row][col+1]
		valid_tile_types := []Tile{EastWest, NorthWest, SouthWest, Start}
		if contains_tile(node.tile, valid_tile_types) {
			neighbors = append(neighbors, [2]int{row, col + 1})
		}
	}
	return neighbors
}

func check_west(maze [][]Node, row, col int, neighbors [][2]int) [][2]int {
	if col > 0 {
		node := maze[row][col-1]
		valid_tile_types := []Tile{EastWest, NorthEast, SouthEast, Start}
		if contains_tile(node.tile, valid_tile_types) {
			neighbors = append(neighbors, [2]int{row, col - 1})
		}
	}
	return neighbors
}

func get_valid_neighbors(maze [][]Node, coord [2]int) (neighbors [][2]int) {
	row, col := coord[0], coord[1]
	curr_node := maze[row][col]
	curr_tile := curr_node.tile

	// All need to check if neighbor is Start
	if curr_tile == NorthSouth {
		// Need to check North and South for any South/North facing tiles
		// North
		neighbors = check_north(maze, row, col, neighbors)

		// South
		neighbors = check_south(maze, row, col, neighbors)
	} else if curr_tile == EastWest {
		// Need to check East and West for any West/East facing tiles
		// East
		neighbors = check_east(maze, row, col, neighbors)

		// West
		neighbors = check_west(maze, row, col, neighbors)
	} else if curr_tile == NorthWest {
		// Need to check North and East for any South/East facing tiles
		// North
		neighbors = check_north(maze, row, col, neighbors)

		// West
		neighbors = check_west(maze, row, col, neighbors)
	} else if curr_tile == NorthEast {
		// Need to check North and West for any South/West facing tiles
		// North
		neighbors = check_north(maze, row, col, neighbors)

		// East
		neighbors = check_east(maze, row, col, neighbors)
	} else if curr_tile == SouthWest {
		// Need to check South and West for any North/East facing tiles
		// South
		neighbors = check_south(maze, row, col, neighbors)

		// West
		neighbors = check_west(maze, row, col, neighbors)
	} else if curr_tile == SouthEast {
		// Need to check South and East for any North/West facing tiles
		// South
		neighbors = check_south(maze, row, col, neighbors)

		// East
		neighbors = check_east(maze, row, col, neighbors)
	} else if curr_tile == Ground {
		// Skip
	} else if curr_tile == Start {
		// Need to find in all four directions
		// North
		neighbors = check_north(maze, row, col, neighbors)

		// South
		neighbors = check_south(maze, row, col, neighbors)

		// East
		neighbors = check_east(maze, row, col, neighbors)

		// West
		neighbors = check_west(maze, row, col, neighbors)
	}
	return
}

func in_array(array [][2]int, v [2]int) bool {
	for _, entry := range array {
		if entry == v {
			return true
		}
	}
	return false
}

func build_loop(maze [][]Node, start [2]int) (coordinates [][2]int) {
	coords := [][2]int{start}

	to_visit := get_valid_neighbors(maze, start)[:1]
    coords = append(coords, to_visit[0])

	i := 0
	for len(to_visit) > 0 {
		curr_pos := to_visit[0]
		to_visit = to_visit[1:]

		curr_neighbors := get_valid_neighbors(maze, curr_pos)
		for _, neighbor := range curr_neighbors {
			visited := in_array(coords, neighbor)
			if !visited {
				coords = append(coords, neighbor)
				to_visit = append(to_visit, neighbor)
			}
		}
		i++
	}
	coordinates = coords
	return
}

func Day_10_Part_1() {
	maze, start := Day_10_parse_input()
	loop := build_loop(maze, start)
	fmt.Println(len(loop) / 2)
}

func print_maze(maze [][]Node, loop [][2]int) {
	for row_i, row := range maze {
		for col_i, col := range row {
			if !in_array(loop, [2]int{row_i, col_i}) {
				col.tile = Ground
				row[col_i] = col
			}
			fmt.Print(col)
		}
		fmt.Println()
		maze[row_i] = row
	}
}

func Day_10_Part_2() {
	maze, start := Day_10_parse_input()

	loop := build_loop(maze, start)
	print_maze(maze, loop)

	// Calculate area of Maze using Polygon area calculation
    // https://www.mathopenref.com/coordpolygonarea.html
	i := 0
	area := 0
	for i < len(loop) {
		coord := loop[i]
		var coord_comp [2]int
		if i == len(loop)-1 {
			coord_comp = loop[0]
		} else {
			coord_comp = loop[i+1]
		}
		value := (coord[0] * coord_comp[1]) - (coord[1] * coord_comp[0])
		area = area + value
		i++
	}
	area = area / 2
	if area < 0 {
		area = area * -1
	}

	// Use that area in Pick's theorem to get number of interior points
    // https://en.wikipedia.org/wiki/Pick%27s_theorem
	// A = i + b/2 - 1
	points := area - len(loop)/2 + 1
	fmt.Println(points)
}
