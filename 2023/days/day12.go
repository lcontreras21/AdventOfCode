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

type Spring int64

const (
	Operational Spring = iota
	Damaged
	Unknown
)

func (spring Spring) String() string {
	switch spring {
	case Operational:
		return "."
	case Damaged:
		return "#"
	default: // Unknown
		return "?"
	}
}

func (record Record) String() string {
	s := ""
	for _, spring := range record.springs {
		s = s + Spring.String(spring)
	}

	s = s + "\t" + fmt.Sprintf("%d", record.sizes)
	return s
}

func (record *Record) extend(count int) Record {
	springs := []Spring{}
	for i := 0; i < count; i++ {
		springs = append(springs, record.springs...)
		if i != count-1 {
			springs = append(springs, Unknown)
		}
	}

	sizes := []int{}
	for i := 0; i < count; i++ {
		sizes = append(sizes, record.sizes...)
	}

	new_record := Record{springs: springs, sizes: sizes}
	return new_record
}

func to_Spring(s string) Spring {
	switch s {
	case ".":
		return Operational
	case "#":
		return Damaged
    default:  // If it is "?"
		return Unknown
	}
}

type Record struct {
	springs []Spring
	sizes   []int
}

func Day_12_parse_input(test int) (records []Record) {
    var filename string
    if test != 0 {
        filename = "2023/inputs/Day_12.txt"
    } else {
        filename = "2023/inputs/temp.txt"
    }
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("[0-9]+")

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, " ")

		matches := r.FindAllString(split[1], -1)
		sizes := []int{}
		for _, m := range matches {
			value, _ := strconv.Atoi(m)
			sizes = append(sizes, value)
		}

		sprongs := strings.Split(split[0], "")
		springs := []Spring{}
		for _, s := range sprongs {
			springs = append(springs, to_Spring(s))
		}

		record := Record{springs: springs, sizes: sizes}
		records = append(records, record)
	}
	file.Close()
	return
}

func print_records(records []Record) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func contains_spring(springs []Spring, spring Spring) bool {
	for _, s := range springs {
		if s == spring {
			return true
		}
	}
	return false
}

func generate_all_possibilities(springs []Spring, sizes []int, number_of_damaged int) int {
	if len(sizes) == 0 {
		return 0
	}

	if number_of_damaged == sizes[0] {
		// Consume a size
		number_of_damaged = -1
		sizes = sizes[1:]
		if len(sizes) == 0 {
			if contains_spring(springs, Damaged) {
				return 0 // Not a valid scenario
			} else {
				return 1 // Found a valid scenario
			}
		}
	}

	if len(springs) == 0 {
		return 0 // Not a valid scenario
	}

	spring := springs[0]
	if spring == Operational {
		if number_of_damaged > 0 {
			return 0 // Not a valid scenario
		} else {
			return generate_all_possibilities(springs[1:], sizes, 0)
		}
	} else if spring == Damaged {
		if number_of_damaged < 0 {
			return 0 // Not a valid scenario
		} else {
			return generate_all_possibilities(springs[1:], sizes, number_of_damaged+1)
		}
	} else { // If spring is Unknown
		scenario_1 := append([]Spring{Operational}, springs[1:]...)
		scenario_2 := append([]Spring{Damaged}, springs[1:]...)

		if_its_operational := generate_all_possibilities(scenario_1, sizes, number_of_damaged)
		if_its_damaged := generate_all_possibilities(scenario_2, sizes, number_of_damaged)
		return if_its_operational + if_its_damaged
	}
}

func Day_12_Part_1() {
	records := Day_12_parse_input(1)

	total := 0
	for _, record := range records {
		count := generate_all_possibilities(record.springs, record.sizes, 0)
		total = total + count
	}
	fmt.Println("count", total)
}

func unfold_records(records []Record) (unfolded []Record) {
	for _, record := range records {
		new_record := record.extend(5)
		unfolded = append(unfolded, new_record)
	}
	return
}

func generate_all_possibilites_with_cache(springs []Spring, sizes []int, cache [][]int) (int) {
    if len(sizes) == 0 {
        if contains_spring(springs, Damaged) {
            return 0
        } else {
            return 1
        }
    }

    if len(springs) < (utils.Sum(sizes) + len(sizes)) {
        return 0
    }

    cache_value := cache[len(sizes) - 1][len(springs) - 1]
    if cache_value >= 0 {
       return cache_value 
    }

    var possibilities int
    if springs[0] != Damaged {
        possibilities = possibilities + generate_all_possibilites_with_cache(springs[1:], sizes, cache)
    }
    next_size := sizes[0]
    if !contains_spring(springs[:next_size], Operational) && springs[next_size] != Damaged {
        possibilities = possibilities + generate_all_possibilites_with_cache(springs[next_size+1:], sizes[1:], cache) 
    }
    cache[len(sizes)-1][len(springs)-1] = possibilities
    return possibilities
}

func Day_12_Part_2() {
    // Adapted from https://nickymeuleman.netlify.app/garden/aoc2023-day12
	records := Day_12_parse_input(1)
	records = unfold_records(records)
    
	total := 0
	for _, record := range records {
        cache := [][]int{}
        record.springs = append(record.springs, Operational)
        for range record.sizes {
            c := []int{}
            for range record.springs {
                c = append(c, -1)
            }
            cache = append(cache, c)
        }
		count := generate_all_possibilites_with_cache(record.springs, record.sizes, cache)
		total = total + count
	}
	fmt.Println("count", total)
}
