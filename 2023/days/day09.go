package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day_9_parse_input() (histories [][]int) {
	file, err := os.Open("inputs/Day_09.txt")
	// file, err := os.Open("inputs/temp.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("[-0-9]+")

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		h := r.FindAllString(txt, -1)
		history := []int{}
		for _, x := range h {
			value, _ := strconv.Atoi(x)
			history = append(history, value)
		}
		histories = append(histories, history)
	}
	file.Close()
	return

}

func extrapolate_backwards(sequences [][]int) (pv int) {
	for i := len(sequences) - 1; i >= 0; i-- {
		sequence := sequences[i]
		pv = sequence[0] - pv
	}
	return
}

func extrapolate_forwards(sequences [][]int) (nv int) {
	for i := len(sequences) - 1; i >= 0; i-- {
		sequence := sequences[i]
		nv = nv + sequence[len(sequence)-1]
	}
	return
}

func Day_9(part int) {
	histories := Day_9_parse_input()

	total := 0
	for _, history := range histories {
		done := 0
		sequences := [][]int{history}
		for done < 1 {
			diff := utils.Diff(history)
			sequences = append(sequences, diff)
			if utils.None(diff) {
				break
			}
			history = diff
		}
		fn := extrapolate_forwards
		if part == 2 {
			fn = extrapolate_backwards
		}

		num := fn(sequences)
		total = total + num
	}
	fmt.Println(total)
}
