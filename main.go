package main

import (
	days_2023 "AdventOfCode/2023/days"
	days_2024 "AdventOfCode/2024/days"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("Please pass in Year Day Part") 
        return
    }

    if len(args) < 2 {
        fmt.Println("Please pass in Day Part") 
        return
    }

    if len(args) < 3 {
        fmt.Println("Please pass in Part") 
        return
    }

	year := args[0]
	day := args[1]
	part := args[2]
	fmt.Println("Working on Year " + year)
	fmt.Println("Working on Day " + day)
	fmt.Println("Working on Part " + part)

    days := days_2024.GetMapping()
    switch year {
    case "2023":
        days = days_2023.GetMapping()
    }

    fn_key := day + "_" + part
    fn, error := days[fn_key]
    if !error {
        fmt.Println("Could not find " + fn_key + " in mapping")
        return
    }
    fn()
}
