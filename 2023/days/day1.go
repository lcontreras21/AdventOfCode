package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
    "AdventOfCode/utils"
)


func Day_1_Part_1() {
	// Read in File
	// file, err := os.Open("Inputs/Day_1.txt")
	file, err := os.Open("Inputs/Day_1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("[0-9]") // Part One

	total := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()

		matches := r.FindAllString(txt, -1)
		first_int := matches[0]
		last_int := matches[len(matches)-1]

		combined_int := first_int + last_int

		to_add, _ := strconv.Atoi(combined_int)
		total = total + to_add
	}

	fmt.Println(total)
	file.Close()
}

func Day_1_Part_2() {
	// Sliding window approach of size 5
	// dhaj2ksj3twone433kthree
	// read in a character, if first character is [0-9] mark as first int, continue otherwise
	// read in a character from right

	file, err := os.Open("Inputs/Day_1.txt")
	// file, err := os.Open("Inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	// Initialize string -> int conversion for Part 2
	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"
	m["four"] = "4"
	m["five"] = "5"
	m["six"] = "6"
	m["seven"] = "7"
	m["eight"] = "8"
	m["nine"] = "9"
	m["1"] = "1"
	m["2"] = "2"
	m["3"] = "3"
	m["4"] = "4"
	m["5"] = "5"
	m["6"] = "6"
	m["7"] = "7"
	m["8"] = "8"
	m["9"] = "9"
	// three, four, five lengths

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		// fmt.Println(txt)

		first_int, second_int := "", ""
		for i := 0; i < len(txt); i++ {
			subtext := txt[i:utils.Min(i+5, len(txt))]
			subtext_3 := subtext[0:utils.Min(3, len(subtext))]
			subtext_4 := subtext[0:utils.Min(4, len(subtext))]

			first_char := subtext[0:1]

			_, prs := m[first_char]
			if prs {
				first_int = first_char
				break
			}
			_, prs = m[subtext_3]
			if prs {
				first_int = m[subtext_3]
				break
			}
			_, prs = m[subtext_4]
			if prs {
				first_int = m[subtext_4]
				break
			}
			_, prs = m[subtext]
			if prs {
				first_int = m[subtext]
				break
			}
		}

		for i := len(txt); i >= 0; i-- {
			subtext := txt[utils.Max(i-5, 0):i]
			subtext_3 := subtext[utils.Max(0, len(subtext)-3):]
			subtext_4 := subtext[utils.Max(0, len(subtext)-4):]

            last_char := subtext[len(subtext)-1:]

			_, prs := m[last_char]
			if prs {
				second_int = last_char
				break
			}
			_, prs = m[subtext_3]
			if prs {
				second_int = m[subtext_3]
				break
			}
			_, prs = m[subtext_4]
			if prs {
				second_int = m[subtext_4]
				break
			}
			_, prs = m[subtext]
			if prs {
				second_int = m[subtext]
				break
			}
		}

		combined_int := first_int + second_int
		// fmt.Println(first_int, second_int)

		to_add, _ := strconv.Atoi(combined_int)
		total = total + to_add
	}

	fmt.Println(total)
	file.Close()

}
