package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func find_in_range(v int, m map[[2]int][2]int) (conv int) {
	conv = v
	for range_s, range_d := range m {
		if v >= range_s[0] && v <= range_s[1] {
            // find offset
            offset := v - range_s[0]
            conv = range_d[0] + offset
		}
	}
	return
}

func find_ranges(ranges [][2]int, m map[[2]int][2]int) (output [][2]int) {
    // keep new ranges
    // go through ranges
    // check conversions
    // if range is within mapping
        // create new range 

    i := 0
    for i < len(ranges) {
        r := ranges[i]
        matched := false
        for range_s, range_d := range m {
            // range 0 9
            // map [2 4]->[4 6], [9 11]->[10 12]
            // os = 2, oe = 4 -> os < oe means overlap
            // add range os-2+4, oe-2+4 -> [4, 6]
            // os = 9, oe = 9
            // add range os-9+10, oe-9+10 -> [10 10]
            // [   ]
            //   []
            //[  ]
            //   [   ]

            // then if some range wasnt covered by map, add those non overlaps to be discovered next
            
            // Get range of potential overlap between given range and mapping range
            overlap_start := utils.Max(r[0], range_s[0])
            overlap_end := utils.Min(r[1], range_s[1])
            if overlap_start <= overlap_end {
                matched = true

                // Add overlapping range to new mapped output
                mapped_start := overlap_start - range_s[0] + range_d[0]
                mapped_end := overlap_end - range_s[1] + range_d[1]
                overlapped_range := [2]int{mapped_start, mapped_end} 
                output = append(output, overlapped_range)

                // Check if some range wasn't covered above
                if overlap_end < r[1] {
                    ranges = append(ranges, [2]int{overlap_end+1, r[1]})
                }

                // Check if some range wasn't covered below
                if overlap_start > r[0] {
                    ranges = append(ranges, [2]int{r[0], overlap_start-1})
                }
                
                // Break since we found a mapping
                break
            }
        }
        if !matched {
            output = append(output, r)
        }
        i++
    }
    return
}

func parse_input_day_5() (seeds []string, maps map[string]map[[2]int][2]int) {
	file, err := os.Open("inputs/Day_5.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	seeds = []string{}

	var doing int
	map_to_doing := map[string]int{
		"seed-to-soil":            1,
		"soil-to-fertilizer":      2,
		"fertilizer-to-water":     3,
		"water-to-light":          4,
		"light-to-temperature":    5,
		"temperature-to-humidity": 6,
		"humidity-to-location":    7,
	}

    maps = make(map[string]map[[2]int][2]int)
	maps["soil_map"] = make(map[[2]int][2]int)
	maps["fertilizer_map"] = make(map[[2]int][2]int)
	maps["water_map"] = make(map[[2]int][2]int)
	maps["light_map"] = make(map[[2]int][2]int)
	maps["temperature_map"] = make(map[[2]int][2]int)
	maps["humidity_map"] = make(map[[2]int][2]int)
	maps["location_map"] = make(map[[2]int][2]int)
	curr_map := maps["soil_map"]

	r, _ := regexp.Compile("[0-9]+")

	for fileScanner.Scan() {
		txt := fileScanner.Text()

		if txt == "" {
			continue
		}

		// Get seeds
		if strings.Contains(txt, "seeds:") {
			seeds_s := strings.Split(txt, ":")[1]
			seeds = r.FindAllString(seeds_s, -1)
			continue
		}

		if strings.Contains(txt, "map") {
			map_name := strings.Split(txt, " ")[0]
			doing = map_to_doing[map_name]
			switch doing {
			case 1:
				curr_map = maps["soil_map"]
			case 2:
				curr_map = maps["fertilizer_map"]
			case 3:
				curr_map = maps["water_map"]
			case 4:
				curr_map = maps["light_map"]
			case 5:
				curr_map = maps["temperature_map"]
			case 6:
				curr_map = maps["humidity_map"]
			case 7:
				curr_map = maps["location_map"]
			}
			continue
		}

		numbers := r.FindAllString(txt, -1)
		length, _ := strconv.Atoi(numbers[2])
		start_s, _ := strconv.Atoi(numbers[1])
		end_s := start_s + length - 1

		start_d, _ := strconv.Atoi(numbers[0])
		end_d := start_d + length - 1

		source := [2]int{start_s, end_s}
		destination := [2]int{start_d, end_d}
		curr_map[source] = destination
	}
	file.Close()
    return
}

func Day_5_Part_1() {
    seeds, maps := parse_input_day_5() 

	// Go through the black box
    min := []int{}
	for _, s := range seeds {
        seed, _ := strconv.Atoi(s)
        soil := find_in_range(seed, maps["soil_map"])
        fertilizer := find_in_range(soil, maps["fertilizer_map"])
        water := find_in_range(fertilizer, maps["water_map"])
        light := find_in_range(water, maps["light_map"])
        temperature := find_in_range(light, maps["temperature_map"])
        humidity:= find_in_range(temperature, maps["humidity_map"])
        location := find_in_range(humidity, maps["location_map"])
        min = append(min, location)
	}
    min_loc := utils.MinArray(min)
    fmt.Println(min_loc)
}

func Day_5_Part_2() {
    seeds, maps := parse_input_day_5() 

	// Go through the black box but less brute-forcey this time
    ranges := [][2]int{}
    i := 0
    for i < len(seeds) {
        s_0, _ := strconv.Atoi(seeds[i])
        s_1, _ := strconv.Atoi(seeds[i+1])
        ranges = append(ranges, [2]int{s_0, s_0 + s_1 - 1})
        i = i + 2
    }

    ranges = find_ranges(ranges, maps["soil_map"])
    ranges = find_ranges(ranges, maps["fertilizer_map"])
    ranges = find_ranges(ranges, maps["water_map"])
    ranges = find_ranges(ranges, maps["light_map"])
    ranges = find_ranges(ranges, maps["temperature_map"])
    ranges = find_ranges(ranges, maps["humidity_map"])
    ranges = find_ranges(ranges, maps["location_map"])
    min := ranges[0][0]
    for _, r := range ranges {
        min = utils.Min(min, r[0])
    }
    fmt.Println(min)
}
