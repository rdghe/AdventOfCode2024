package day1

import (
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(part string, example bool) error {
	if part != "1" && part != "2" {
		return errors.New("invalid part specified; it should be either `1` or `2`")
	}

	log.Printf("solving day 1 part %s\n", part)

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
