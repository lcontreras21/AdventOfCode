package days

import (
	"bufio"
	"fmt"
	"os"
)


func Day_1_Parse_Input() {
	// Read in File
	file, err := os.Open("inputs/FILENAME.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
        fmt.Println(txt)
	}

	fmt.Println(total)
	file.Close()
}

func Day_1_Part_1() {
}

func Day_1_Part_2() {
}
