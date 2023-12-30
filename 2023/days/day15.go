package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	label        string
	focal_length int
}

type Box struct {
	lens_slots []Lens
}

func Day_15_parse_input(use_test_file bool) (sequences [][]string) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_15.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	line := []string{}
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		line = strings.Split(txt, ",")
	}
	for _, l := range line {
		sequences = append(sequences, strings.Split(l, ""))
	}

	file.Close()
	return
}

func run_hash_algorithm(label []string) (current_value int) {
	for _, letter := range label {
		ascii_code := int(letter[0])
		current_value = current_value + ascii_code
		current_value = current_value * 17
		current_value = current_value % 256
	}
	return
}

func Day_15_Part_1() {
	// sequences := Day_15_parse_input(true)
	sequences := Day_15_parse_input(false)

	total := 0
	for _, sequence := range sequences {
		current_value := run_hash_algorithm(sequence)
		total = total + current_value
	}
	fmt.Println(total)
}

func print_hashmap(hashmap [256]Box) {
	for i, box := range hashmap {
		if len(box.lens_slots) > 0 {
			fmt.Println("Box", i, box.lens_slots)
		}
	}
}

func hashmapify(sequences [][]string) (hashmap [256]Box) {
	for _, sequence := range sequences {
		var box_number int
		var label, operation, label_placeholder string
		for _, letter := range sequence {
			if letter == "=" || letter == "-" {
				operation = letter
				label = label_placeholder
				box_number = run_hash_algorithm(strings.Split(label, ""))
			}
			label_placeholder = label_placeholder + letter
		}

		box := hashmap[box_number]
		if operation == "=" {
			found := false
			focal_length, _ := strconv.Atoi(sequence[len(sequence)-1])
			for i, lens := range box.lens_slots {
				if lens.label == label {
					found = true
					box.lens_slots[i].focal_length = focal_length
					break
				}
			}
			if !found {
				new_lens := Lens{focal_length: focal_length, label: label}
				box.lens_slots = append(box.lens_slots, new_lens)
			}
		} else {
			for i, lens := range box.lens_slots {
				if lens.label == label {
					box.lens_slots = append(box.lens_slots[:i], box.lens_slots[i+1:]...)
					break
				}
			}
		}
		hashmap[box_number] = box
	}

	return
}

func Day_15_Part_2() {
	// sequences := Day_15_parse_input(true)
	sequences := Day_15_parse_input(false)

	hashmap := hashmapify(sequences)
	// print_hashmap(hashmap)

	total := 0
	for i, box := range hashmap {
		box_number := i + 1
        for j, lens := range box.lens_slots {
            slot := j + 1
            focusing_power := box_number * slot * lens.focal_length
            total = total + focusing_power
        }
	}
    fmt.Println(total)
}
