package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// parse each line into direction and clicks
	scanner := bufio.NewScanner(file)

	totalTimesDialLandsOnZeros := 0
	// part 2 (gold star challenge):
	totalZeroClicks := 0

	// starting positions
	position := 50

	for scanner.Scan() {
		line := scanner.Text()

		direction := string(line[0])
		clicksStr := line[1:]

		var clicks int
		fmt.Sscanf(clicksStr, "%d", &clicks)

		fmt.Printf("direction: %s, clicks: %d\n", direction, clicks)

		fullRotations := clicks / 100
		remainder := clicks % 100

		totalZeroClicks += fullRotations

		if direction == "R" {
			newPosition := position + remainder
			if newPosition > 100 {
				// means we crossed (and clicked) 0 in the CW direction
				totalZeroClicks++
			}

			position = newPosition % 100
			if position == 0 {
				totalTimesDialLandsOnZeros++
			}

		}
		if direction == "L" {
			newPosition := position - remainder
			if position != 0 && newPosition < 0 {
				// means we crossed (and clicked) 0 in the CW direction
				totalZeroClicks++
			}

			position = (newPosition%100 + 100) % 100
			if position == 0 {
				totalTimesDialLandsOnZeros++
			}
		}

		fmt.Printf("current positions: %d\n", position)
		fmt.Printf("total time dial lands on zero: %d\n", totalTimesDialLandsOnZeros)
		fmt.Printf("total zero clicks: %d\n", totalZeroClicks)
		fmt.Println("======")
	}

	fmt.Printf("TOTAL: %d\n", totalTimesDialLandsOnZeros+totalZeroClicks)

	// any errors during scanning?
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}

}
