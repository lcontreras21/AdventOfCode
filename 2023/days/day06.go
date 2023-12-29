package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse_input_day_6() (output map[int]int, times, distances []string) {
	file, err := os.Open("inputs/Day_06.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	line_1 := ""
	line_2 := ""
	i := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if i == 0 {
			line_1 = txt
		} else {
			line_2 = txt
		}
		i++
	}
	file.Close()

	r, _ := regexp.Compile("[0-9]+")
	times = r.FindAllString(line_1, -1)
	distances = r.FindAllString(line_2, -1)

	output = make(map[int]int)
	i = 0
	for i < len(times) {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		output[time] = distance
		i++
	}
	return
}

func Day_6_Part_1() {
	races, _, _ := parse_input_day_6()

	total := 1
	for time, record := range races {
		t := 1
		ways_beat := 0
		for t < time {
			if (time*t - utils.Power(t, 2)) > record {
				ways_beat++
			}
			t++
		}
		total = total * ways_beat
	}
	fmt.Println(total)
}

func Day_6_Part_2() {
    // Part 1 uses brute force
    // This one uses quadratic formula wooooooh
	_, times, distances := parse_input_day_6()
	time := strings.Join(times, "")
	distance := strings.Join(distances, "")

    t, _ := strconv.Atoi(time)
    d, _ := strconv.Atoi(distance)

	l, r := utils.QuadraticFormula(-1, t, -1*d)
	left := int(math.Ceil(l))
	right := int(math.Floor(r))

    // Add one if the intercepts are actually integer values because that means
    // we'd match the record
    if -1 * utils.Power(left, 2) + left * t - d == 0 {
        left++
    }
    if -1 * utils.Power(right, 2) + right * t - d == 0 {
        right++
    }
    diff := right - left + 1
    fmt.Println(diff)
}
