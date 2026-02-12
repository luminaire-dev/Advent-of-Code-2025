package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bank := make([]int, len(line))
		for i, ch := range line {
			bank[i], _ = strconv.Atoi(string(ch))
		}

		jolts := 0
		start := 0
		for index := 0; index < 11 && start < len(bank); index++ {
			maxDigit := -1
			maxIdx := -1
			for i := start; i < len(bank)-(11-index); i++ {
				if bank[i] > maxDigit {
					maxDigit = bank[i]
					maxIdx = i
				}
			}
			if maxIdx == -1 {
				break
			}
			jolts = jolts*10 + maxDigit
			start = maxIdx + 1
		}
		// Add the last max digit if any left
		if start < len(bank) {
			maxDigit := bank[start]
			for i := start; i < len(bank); i++ {
				if bank[i] > maxDigit {
					maxDigit = bank[i]
				}
			}
			jolts = jolts*10 + maxDigit
		}
		total += jolts
	}
	fmt.Println(total)
}
