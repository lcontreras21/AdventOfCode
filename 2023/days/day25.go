package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day_25_parse_input(use_test_file bool) ()  {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_25.txt"
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
		// txt := fileScanner.Text()
	}

	file.Close()
	return

}

func Day_25_Part_1() {
}
