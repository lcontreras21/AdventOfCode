package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day_24_parse_input(use_test_file bool) ()  {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_24.txt"
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

func Day_24_Part_1() {
}
