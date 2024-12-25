package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solve(part2 bool, example bool) error {
	var filename string
	if example {
		filename = "inputs/day3_example.txt"
	} else {
		filename = "inputs/day3.txt"
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
	log.Println("solving day 3 part 1")
	result := 0

	log.Println(inputString)
	mulRegex, _ := regexp.Compile(`mul\((?P<LeftNum>\d+),(?P<RightNum>\d+)\)`)

	// Find all submatches of the string which uses capture groups to extract the two elements
	matches := mulRegex.FindAllStringSubmatch(inputString, -1)
	log.Printf("found the following matches: %v\n", matches)

	for _, match := range matches {
		log.Println(match)
		log.Println(match[1])
		log.Println(match[2])
		left, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		}
		right, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		result += left * right
	}

	log.Printf("result is: %d\n", result)
	return nil
}

func solvePartTwo(inputString string) error {
	log.Println("solving day 3 part 2")
	result := 0

	log.Println(inputString)
	//mulRegex, _ := regexp.Compile(`mul\((?P<LeftNum>\d+),(?P<RightNum>\d+)\)`)
	doRegex, _ := regexp.Compile(`do\(\)`)
	dontRegex, _ := regexp.Compile(`don't\(\)`)

	dos := doRegex.FindStringIndex(inputString)
	log.Println(dos)
	donts := dontRegex.FindStringIndex(inputString)
	log.Println(donts)

	//// Find all submatches of the string which uses capture groups to extract the two elements
	//matches := mulRegex.FindAllStringSubmatch(inputString, -1)
	//log.Printf("found the following matches: %v\n", matches)
	//
	//for _, match := range matches {
	//	log.Println(match)
	//	log.Println(match[1])
	//	log.Println(match[2])
	//	left, err := strconv.Atoi(match[1])
	//	if err != nil {
	//		return err
	//	}
	//	right, err := strconv.Atoi(match[2])
	//	if err != nil {
	//		return err
	//	}
	//
	//	result += left * right
	//}

	log.Printf("result is: %d\n", result)
	return nil
}
