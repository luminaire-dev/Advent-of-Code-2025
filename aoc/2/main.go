package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read csv file
	f, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	invalidIDsAddedUp := 0

	// read each csv range from the file
	for scanner.Scan() {
		line := scanner.Text()
		idRanges := strings.Split(line, ",")
		for _, r := range idRanges {

			// split on - to seperating min and max values from range
			rangeVals := strings.Split(r, "-")

			fmt.Print("====\n")
			fmt.Printf("min: %s \n", rangeVals[0])
			fmt.Printf("max: %s \n", rangeVals[1])

			minInt, err := strconv.Atoi(rangeVals[0])
			if err != nil {
				panic(err)
			}
			maxInt, err := strconv.Atoi(rangeVals[1])
			if err != nil {
				panic(err)
			}

			for i := minInt; i <= maxInt; i++ {
				numberStr := strconv.Itoa(i)

				if len(numberStr) == 1 {
					continue
				}

				// if all characters in numberStr are the same increment invalid IDs and continue to next number
				allSame := strings.Count(numberStr, string(numberStr[0])) == len(numberStr)
				if allSame {
					fmt.Printf("%d contains all identical digits\n", i)
					invalidIDsAddedUp += i
					continue
				}

				// check for repeated patterns of any length, from 2 up to half the length of the string.
				// (half because if the pattern > than half, it cannot be repeating)
				for patternLen := 2; patternLen <= len(numberStr)/2; patternLen++ {
					if len(numberStr)%patternLen != 0 {
						continue
					}
					pattern := numberStr[:patternLen]
					repeats := len(numberStr) / patternLen
					expected := strings.Repeat(pattern, repeats)
					if numberStr == expected {
						fmt.Printf("%d contains repeated pattern %s, %d times\n", i, pattern, repeats)
						invalidIDsAddedUp += i
						break
					}
				}
			}
		}

		fmt.Println(invalidIDsAddedUp)

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

// part one:

// func main() {

// 	// read csv file
// 	f, err := os.Open("input.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	invalidIDsAddedUp := 0

// 	// read each csv range from the file
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		idRanges := strings.Split(line, ",")
// 		for _, r := range idRanges {

// 			// split on - to seperating min and max values from range
// 			rangeVals := strings.Split(r, "-")

// 			fmt.Print("====\n")
// 			fmt.Printf("min: %s \n", rangeVals[0])
// 			fmt.Printf("max: %s \n", rangeVals[1])

// 			minInt, err := strconv.Atoi(rangeVals[0])
// 			if err != nil {
// 				panic(err)
// 			}
// 			maxInt, err := strconv.Atoi(rangeVals[1])
// 			if err != nil {
// 				panic(err)
// 			}

// 			for i := minInt; i <= maxInt; i++ {

// 				numberString := strconv.Itoa(i)
// 				if len(numberString)%2 != 0 {
// 					// ignore numbers that cannot be split into two equal parts,
// 					// meaning they cannot contain repeating sequences as described in the problem
// 					fmt.Printf("ignoring %d, it cannot be split into two equal parts \n", i)
// 					continue
// 				}

// 				half := len(numberString) / 2
// 				head := numberString[0:half]
// 				tail := numberString[half:]

// 				if head == tail {
// 					fmt.Printf("%d contains repeating sequence \n", i)
// 					invalidIDsAddedUp += i
// 				}
// 			}

// 			fmt.Println(invalidIDsAddedUp)

// 		}

// 	}

// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}

// }
