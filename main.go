package main

import (
	"fmt"
	"os"
)

func main() {
	problem := os.Args[1]
	var fileName string
	if len(os.Args) > 2 {
		fileName = os.Args[2]
	} else {
		fileName = fmt.Sprintf("input/day%s.txt", problem)
	}
	fmt.Printf("Running problem: %s with file: %s\n", problem, fileName)
	input := ReadLines(fileName)
	switch problem {
	case "01":
		Day01(input)
	case "02":
		Day02(input)
	case "03":
		Day03(input)
	case "04":
		Day04(input)
	case "05":
		Day05(input)
	case "06":
		Day06(input)
	case "07":
		Day07(input)
	case "08":
		Day08(input)
	case "09":
		Day09(input)
	case "10":
		Day10(input)
	case "11":
		Day11(input)
	case "12":
		Day12(input)
	case "19":
		Day19(input)
	default:
		fmt.Printf("Problem %s not implemented\n", problem)
	}

}
