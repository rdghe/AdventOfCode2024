package day1

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(part2 bool, example bool) error {
	var filename string
	if example {
		filename = "inputs/day1_example.txt"
	} else {
		filename = "inputs/day1.txt"
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
	lines := strings.Split(inputString, "\n")
	var left []int
	var right []int

	for _, line := range lines {
		// Split each line into fields
		fields := strings.Fields(line)

		if len(fields) == 2 {
			// Convert the fields to integers and append them to the respective slices
			leftNum, err := strconv.Atoi(fields[0])
			if err != nil {
				return err
			}
			rightNum, err := strconv.Atoi(fields[1])
			if err != nil {
				return err
			}
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	result := 0

	for idx, _ := range left {
		if left[idx] < right[idx] {
			result += right[idx] - left[idx]
		} else {
			result += left[idx] - right[idx]
		}
	}

	log.Printf("the result is: %d\n", result)

	return nil
}

func solvePartTwo(inputString string) error {
	log.Println("solving day 1 part 2")

	lines := strings.Split(inputString, "\n")
	var left []int
	var right []int

	for _, line := range lines {
		// Split each line into fields
		fields := strings.Fields(line)

		if len(fields) == 2 {
			// Convert the fields to integers and append them to the respective slices
			leftNum, err := strconv.Atoi(fields[0])
			if err != nil {
				return err
			}
			rightNum, err := strconv.Atoi(fields[1])
			if err != nil {
				return err
			}
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	similarityScore := 0
	for _, leftNumber := range left {
		appearances := 0
		for _, rightNumber := range right {
			if leftNumber == rightNumber {
				appearances++
			}
		}
		similarityScore += leftNumber * appearances
	}

	log.Printf("the result is: %d\n", similarityScore)

	return nil
}
