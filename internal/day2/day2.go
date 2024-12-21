package day2

import (
	"log"
	"os"
)

func Solve(part2 bool, example bool) error {
	var filename string
	if example {
		filename = "inputs/day2_example.txt"
	} else {
		filename = "inputs/day2.txt"
	}

	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	inputString := string(inputBytes)

	if part2 {
		err := solvePartTwo(inputString)
		if err != nil {
			return err
		}
	} else {
		err := solvePartOne(inputString)
		if err != nil {
			return err
		}
	}

	return nil
}

func solvePartOne(inputString string) error {
	log.Println("solving day 1 part 1")

	return nil
}

func solvePartTwo(inputString string) error {
	log.Println("solving day 1 part 2")

	return nil
}
