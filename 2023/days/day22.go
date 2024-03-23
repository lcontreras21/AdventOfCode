package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
	z int
}

func (c Coord) String() string {
	return fmt.Sprintf("[%d, %d, %d]", c.x, c.y, c.z)
}

type Brick struct {
	start   Coord
	end     Coord
	settled bool
}

func (b Brick) String() string {
	return Coord.String(b.start) + " - " + Coord.String(b.end)
}

func (b *Brick) can_fall(occupied map[Coord]int) bool {
	// Ground at z = 0
	above_ground := b.start.z > 1

	x_range := utils.RangeInclusive(b.start.x, b.end.x, 1)
	y_range := utils.RangeInclusive(b.start.y, b.end.y, 1)

	c := utils.CartesianProduct[int](x_range, y_range)

	is_occupied := false
	for obj := range c {
		pos := Coord{x: obj[0], y: obj[1], z: b.start.z - 1}
		_, contained := occupied[pos]
		if contained {
			is_occupied = contained
			break
		}
	}

	return above_ground && !is_occupied
}

func (b *Brick) fall(occupied map[Coord]int) {
	for b.can_fall(occupied) {
		b.start.z -= 1
		b.end.z -= 1
	}
}

func brick_from_loc(start_loc, end_loc [3]int) Brick {
	start_coord := Coord{x: start_loc[0], y: start_loc[1], z: start_loc[2]}
	end_coord := Coord{x: end_loc[0], y: end_loc[1], z: end_loc[2]}
	new_brick := Brick{start: start_coord, end: end_coord}
	return new_brick
}

func Day_22_parse_input(use_test_file bool) (bricks []Brick) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_22.txt"
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
		split := strings.Split(txt, "~")
		start_txt, end_txt := split[0], split[1]

		start_pos_str := strings.Split(start_txt, ",")
		end_pos_str := strings.Split(end_txt, ",")

		var start_pos, end_pos [3]int
		for i, v := range start_pos_str {
			v_int, _ := strconv.Atoi(v)
			start_pos[i] = v_int
		}
		for i, v := range end_pos_str {
			v_int, _ := strconv.Atoi(v)
			end_pos[i] = v_int
		}
		bricks = append(bricks, brick_from_loc(start_pos, end_pos))
	}

	file.Close()
	return
}

func sort_bricks(bricks []Brick) []Brick {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})

	return bricks
}

func settle_bricks(bricks []Brick) (settled_bricks []Brick, above, below map[int]map[int]bool) {
	// Iterate through bricks, skipping settled
	//  Lower z value until we find a intercept another brick in x, y, or z
	//  Or we encounter z = 1.
	//      If (a) and the intercepted brick is not settled then add it to queue
	//      If (b) mark it as settled
	bricks = sort_bricks(bricks)
	occupied := map[Coord]int{}
	above = map[int]map[int]bool{}
	below = map[int]map[int]bool{}

	for i, brick := range bricks {
		brick.fall(occupied)

		// Add in every Coord that the brick occupies to the occupied map
		x_range := utils.RangeInclusive(brick.start.x, brick.end.x, 1)
		y_range := utils.RangeInclusive(brick.start.y, brick.end.y, 1)
		z_range := utils.RangeInclusive(brick.start.z, brick.end.z, 1)

		c := utils.CartesianProduct(x_range, y_range)
		for obj := range c {
			for _, z := range z_range {
				coord := Coord{x: obj[0], y: obj[1], z: z}
				occupied[coord] = i
			}

			below_coord := Coord{x: obj[0], y: obj[1], z: brick.start.z - 1}
			below_index, is_occupied := occupied[below_coord]
			if is_occupied {
				above_set, above_exists := above[below_index]
				if !above_exists {
					above_set = map[int]bool{}
				}
				above_set[i] = true
				above[below_index] = above_set

				below_set, below_exists := below[i]
				if !below_exists {
					below_set = map[int]bool{}
				}
				below_set[below_index] = true
				below[i] = below_set
			}
		}

		// Update brick
		bricks[i] = brick
	}
	settled_bricks = bricks

	return
}

func disintegrate_bricks(bricks []Brick, above, below map[int]map[int]bool) (disintegrated_bricks int) {
	for i := range bricks {
		above_set, above_exists := above[i]
		if above_exists {
			// Check if each brick above it can be supported by another brick
			//  can be disintegrated if all supported
			can_disintegrate := true
			for above_id := range above_set {
				below_set, _ := below[above_id]
				if len(below_set) <= 1 {
					can_disintegrate = false
				}
			}
			if can_disintegrate {
				disintegrated_bricks++
			}
		} else {
			// Can be disintegrated
			disintegrated_bricks++
		}
	}

	return
}

func helper(id int, above, below map[int]map[int]bool, disintegrated map[int]bool) (int) {
    disintegrated[id] = true

    above_set, above_exists := above[id]
    if above_exists {
        for above_id := range(above_set) {
            below_disintegrated_count := 0
            below_set := below[above_id]

            for below_id := range(below_set) {
                _, already_disintegrated := disintegrated[below_id]
                if already_disintegrated {
                    below_disintegrated_count++
                }
            }

            if below_disintegrated_count == len(below_set) {
                helper(above_id, above, below, disintegrated)
            }
        }
    }

    return len(disintegrated) - 1
}

func disintegrate_bricks_fast(bricks []Brick, above, below map[int]map[int]bool) (disintegrated_bricks int) {
    // Iterate through bricks
    // "disintegrate brick by addding it to set of removed bricks
    // iterate through bricks above it
    //  for each above brick, iterate through its below bricks
    //      if they've all been disintegrated, then above brick can be disintegrated
    //      skip if it can't

	for i := range bricks {
        disintegrated_bricks = disintegrated_bricks + helper(i, above, below, map[int]bool{})
	}

	return
}

func Day_22_Part_1() {
	// bricks := Day_22_parse_input(true)
	bricks := Day_22_parse_input(false)

	bricks, above_map, below_map := settle_bricks(bricks)
	remaining_bricks := disintegrate_bricks(bricks, above_map, below_map)

	fmt.Println("Can safely disintegrate", remaining_bricks, "bricks")
}

func Day_22_Part_2() {
	// bricks := Day_22_parse_input(true)
	bricks := Day_22_parse_input(false)

	bricks, above_map, below_map := settle_bricks(bricks)
	total := disintegrate_bricks_fast(bricks, above_map, below_map)

	fmt.Println("Can safely disintegrate", total, "bricks")
}
