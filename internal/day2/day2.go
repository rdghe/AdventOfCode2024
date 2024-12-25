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

func isSafe(arr []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if diff < 1 || diff > 3 {
			isIncreasing = false
		}
		if diff > -1 || diff < -3 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
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

		if isSafe(numbers) {
			log.Printf("is safe")
			safeReports++
		} else {
			log.Printf("is unsafe")
		}
	}

	log.Printf("the result is: %d\n", safeReports)

	return nil
}

func solvePartTwo(inputString string) error {
	log.Println("solving day 2 part 2")
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

		if isSafe(numbers) {
			log.Println("is safe without modification")
			safeReports++
			continue
		}

		// Check if removing one level makes it safe
		isSafeWithRemoval := false
		for i := 0; i < len(numbers); i++ {
			log.Printf("removing %v with index %v", numbers[i], i)
			// Create a new array with the i-th level removed
			modified := make([]int, len(numbers)-1)
			copy(modified, numbers[:i])
			copy(modified[i:], numbers[i+1:])
			log.Printf("created %v", modified)

			if isSafe(modified) {
				log.Printf("becomes safe after removing %d\n", numbers[i])
				isSafeWithRemoval = true
				break
			}
		}

		if isSafeWithRemoval {
			safeReports++
		} else {
			log.Println("is unsafe even with removal")
		}
	}

	log.Printf("the result is: %d\n", safeReports)

	return nil
}
