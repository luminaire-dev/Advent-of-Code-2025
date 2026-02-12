package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parse each line into direction and clicks
	scanner := bufio.NewScanner(file)

	var min []int
	var max []int
	freshCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // ignore empty or whitespace-only lines
		}
		fmt.Println(line)

		if strings.Contains(line, "-") {
			// it's a range input
			rng := strings.Split(line, "-")
			minValue, err := strconv.Atoi(rng[0])
			if err != nil {
				panic(err)
			}
			maxValue, err := strconv.Atoi(rng[1])
			if err != nil {
				panic(err)
			}

			min = append(min, minValue)
			max = append(max, maxValue)

			fmt.Printf("range min is %d \n", minValue)
			fmt.Printf("range max is %d \n", maxValue)

		} else {
			// it's an ID input
			ID, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			//  min and max slices are of same size
			for i := 0; i < len(min); i++ {
				if ID >= min[i] && ID <= max[i] {
					fmt.Printf("id %d falls in range %d - %d \n", ID, min[i], max[i])
					freshCount++
					break
				}
			}
		}
	}

	fmt.Print(freshCount)

}
