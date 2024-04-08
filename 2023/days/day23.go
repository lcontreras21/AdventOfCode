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

type Trail struct {
	loc             models.Coord
	movement        int
	hashed_path     []string
	path            []models.Coord // Would be nice to change this to a set, since no overlap
    dir             Bearing
}

func Day_23_parse_input(use_test_file bool) (hiking_map models.Matrix[string], start, end models.Coord)  {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_23.txt"
	} else {
		filename = "inputs/temp.txt"
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
		line := strings.Split(txt, "")
        hiking_map.AddRow(hiking_map.Rows(), line)
        i++
	}
    i--
    
    start_index := utils.FindIndex(hiking_map.GetRow(0), ".")
    end_index := utils.FindIndex(hiking_map.GetRow(i), ".")
    start = models.Coord{X: 0, Y: start_index}
    end = models.Coord{X: i, Y: end_index}
 	file.Close()
	return
}


func move_trail(trail Trail, dir Bearing) Trail {
	switch dir {
	case North:
        trail.loc = models.Coord{X:trail.loc.X - 1, Y:trail.loc.Y}
	case South:
        trail.loc = models.Coord{X:trail.loc.X + 1, Y:trail.loc.Y}
	case East:
        trail.loc = models.Coord{X:trail.loc.X, Y:trail.loc.Y + 1}
	case West:
        trail.loc = models.Coord{X:trail.loc.X, Y:trail.loc.Y - 1}
	}

	return trail
}

func hashify_trail(trail Trail) string {
	k := strconv.Itoa(trail.loc.X) + "," + strconv.Itoa(trail.loc.Y)
	// k = k + "," + fmt.Sprint(trail.dir)
	// k = k + "," + strconv.Itoa(trail.movement)
	return k
}

func get_possible_trails(hiking_map models.Matrix[string], trail Trail, is_part_two bool) (moved []Trail) {
    hashed := hashify_trail(trail)
    slope_to_bearning := map[string]Bearing{">": East, "<": West, "v": South}

	for _, bearing := range []Bearing{North, South, East, West} {
		updated := move_trail(trail, bearing)
        curr_tile := hiking_map.Get(updated.loc.X, updated.loc.Y)
        prev_tile := hiking_map.Get(trail.loc.X, trail.loc.Y)

        if curr_tile == "#" {
            continue
        }

        if prev_tile == "#" && is_part_two {
            continue
        }

		if !is_valid_loc([2]int{updated.loc.X, updated.loc.Y}, hiking_map.Cols(), hiking_map.Rows()) {
			// Can't move out of bounds
			continue
		}

        if utils.FindIndex(trail.hashed_path, hashed) >= 0 {
            // Cant repeat steps
            continue
        }

        slope_bearing, exists := slope_to_bearning[curr_tile]
        if exists && bearing != slope_bearing && !is_part_two {
            // Can't go in any other direction that isn't in the same dir as the slope
            continue
        }

		// Update movement
        updated.movement++

		// Update Dir
		updated.dir = bearing

		// Update Path
		updated.hashed_path = utils.Clone(updated.hashed_path)
		updated.hashed_path = append(updated.hashed_path, hashed)
		updated.path = utils.Clone(updated.path)
		updated.path  = append(updated.path, updated.loc)

		moved = append(moved, updated)
	}

	return

}

func add_max_queue(trails []Trail, new Trail) []Trail {
	i := 0
	for _, trail := range trails {
		if new.movement > trail.movement {
			break
		}
		i++
	}

	// Insert at index i
	if i == len(trails) {
		trails = append(trails, new)
	} else {
		trails = append(trails[:i+1], trails[i:]...)
		trails[i] = new
	}

	return trails
}

func find_longest_trail(hiking_map models.Matrix[string], start, end models.Coord, is_part_two bool) int {

    start_trail := Trail{loc: start, dir: South}
	trails := []Trail{start_trail}

    completed_trails := []Trail{}

	for len(trails) > 0 {
		trail := trails[0]
		trails = trails[1:]

		next := get_possible_trails(hiking_map, trail, is_part_two)
		for _, next_trail := range next {
			if next_trail.loc.X == end.X && next_trail.loc.Y == end.Y {
                cloned := hiking_map.Clone()
                for _, loc := range(next_trail.path) {
                    cloned.Set(loc.X, loc.Y, "0")
                }
                fmt.Println("Found trail", trail.movement)
                completed_trails = add_max_queue(completed_trails, next_trail)
                continue
                // return next_trail.movement
			}

            trails = add_max_queue(trails, next_trail)
		}
	}

	return completed_trails[0].movement
}

func Day_23_Part_1() {
	// hiking_map, start, end := Day_23_parse_input(true)
	hiking_map, start, end := Day_23_parse_input(false)

    // fmt.Println(hiking_map, start, end)
    path_len := find_longest_trail(hiking_map, start, end, false)
    fmt.Println("Longest trail has length", path_len)
}

func get_forks(hiking_map models.Matrix[string], start, end models.Coord) (forks map[models.Coord]bool) {
    x_range := utils.RangeInclusive(0, hiking_map.Rows(), 1)
    y_range := utils.RangeInclusive(0, hiking_map.Cols(), 1)
	c := utils.CartesianProduct(x_range, y_range)

    forks = map[models.Coord]bool{start: true, end: true}

	for obj := range c {
		pos := models.Coord{X: obj[0], Y: obj[1]}
        trail := Trail{loc: pos}
        neighbors := get_possible_trails(hiking_map, trail, true)
        if len(neighbors) > 2 {
            forks[pos] = true
        }
    }
    return
}

func get_costmap(hiking_map models.Matrix[string], forks map[models.Coord]bool) (costmap map[models.Coord]map[models.Coord]int) {
    costmap = map[models.Coord]map[models.Coord]int{}
    // Keep track of the length of one point to every other point

    // Iterate through each point
    //  have a list of reachable_points to iterate through
    //  iterate through each of those RP
    //      if RP is one of the Points, add cost to costmap
    //      else iterate through neighbors
    //          if neighbor not seen yet, add to RP with added cost
    //      add current RP to seen

    for point := range(forks) {
        next := []Trail{}
        seen := models.Set[models.Coord]{}

        trail := Trail{loc: point, movement: 0, path: []models.Coord{point}}
        next = append(next, trail)

        for len(next) > 0 {
            n := next[0]
            next = next[1:]

            _, is_fork := forks[n.loc]
            if is_fork && n.movement != 0 {
                // Add to costmap
				map_set, mapset_exists := costmap[point]
				if !mapset_exists {
					map_set = map[models.Coord]int{}
				}
				map_set[n.loc] = n.movement
                costmap[point] = map_set
                continue
            }

            // Iterate through neighbors
            neighbors := get_possible_trails(hiking_map, n, true)
            for _, neighbor := range neighbors {
                if !seen.Contains(neighbor.loc) {
                    next = append(next, neighbor)
                }
                seen.Append(neighbor.loc)
            }
            seen.Append(n.loc)
        }
    }
    return
}

func find_longest_trail_part_two(start, end models.Coord,  costmap map[models.Coord]map[models.Coord]int) (int, []models.Coord) {
    // starting with start find longest path from one point to any other
    //  if updated pos is end, update max
    //  otherwise go through costmap entries for current point
    //      if any not currently seen, update movement, add to seen and keep going

    start_trail := Trail{loc: start, path: []models.Coord{start}}
    trails := []Trail{start_trail}

    max := 0
    max_path := []models.Coord{}

    for len(trails) > 0 {
        trail := trails[0]
        trails = trails[1:]

        if trail.loc.X == end.X && trail.loc.Y == end.Y {
            if trail.movement > max {
                max = trail.movement
                max_path = trail.path
            }
            continue
        }

        map_set := costmap[trail.loc]
        for n, distance := range(map_set) {
            if utils.FindIndex(trail.path, n) >= 1 {
                continue
            }
            path := utils.Clone(trail.path)
            path = append(path, n)
            new_trail := Trail{loc: n, movement: distance + trail.movement, path: path}
            trails = append(trails, new_trail)
        }
    }
    return max, max_path
}

func Day_23_Part_2() {
	// hiking_map, start, end := Day_23_parse_input(true)
	hiking_map, start, end := Day_23_parse_input(false)

    // Using the fn find_longest_trail and changing how it calculates slope movement
    // for part two will cause it to run for 1hr+ (I stopped it after 1 hr).
    // Do not try again :(
    // path_len := find_longest_trail(hiking_map, start, end, true)
    // fmt.Println("Longest trail has length", path_len)


    // Get the list of all interesting points (ie forks, start and end)
    // Calculate how many steps it takes to go from one point to any other
    // Use Points and Costs to calculate the longest possible path
    forks := get_forks(hiking_map, start, end)
    cost_map :=  get_costmap(hiking_map, forks)

    // Running it takes a couple seconds but finishes in the end :)
    path_len, _ := find_longest_trail_part_two(start, end, cost_map)
    fmt.Println("Longest trail has length", path_len)
}
