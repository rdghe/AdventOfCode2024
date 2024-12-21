package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
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
	log.Println("solving day 2 part 1")
	lines := strings.Split(inputString, "\n")

	safeReports := 0

	for _, line := range lines {
		var numbers []int
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		for _, i := range fields {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, j)
		}
		log.Println(numbers)

		isSafe := customIsIncreasing(numbers) || customIsDecreasing(numbers)
		if isSafe {
			log.Printf("is safe (increasing: %v, decreasing: %v)", customIsIncreasing(numbers), customIsDecreasing(numbers))
			safeReports++
		} else {
			log.Printf("is unsafe")
		}
	}

	log.Printf("the result is: %d\n", safeReports)

	return nil
}

func customIsIncreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		previous := arr[i-1]
		difference := current - previous
		if current <= previous || difference > 3 {
			return false
		}
	}
	return true
}

func customIsDecreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		previous := arr[i-1]
		difference := current - previous
		if current >= previous || difference < -3 {
			return false
		}
	}
	return true
}

func solvePartTwo(inputString string) error {
	log.Println("solving day 2 part 2")

	return nil
}
