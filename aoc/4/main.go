package main

import (
	"fmt"
	"os"
	"strings"
)

// part 2:
func main() {
	directions := [][]int{
		//x, y
		{-1, -1}, // bottom-left
		{-1, 0},  // left
		{-1, 1},  // top-left
		{0, 1},   // top
		{1, 1},   // top-right
		{1, 0},   // right
		{1, -1},  // right-bottom
		{0, -1},  // bottom
	}
	fmt.Print(directions)
	fmt.Print(" \n")

	inputBytes, _ := os.ReadFile("input2.txt")
	lines := strings.Split(strings.TrimSpace(string(inputBytes)), "\n")

	fmt.Print(lines)

	rowSize := len(lines)
	colSize := len(lines[0])

	grid := make([][]string, rowSize)
	for i, line := range lines {
		grid[i] = make([]string, colSize)
		for j, val := range line {
			grid[i][j] = string(val)
		}
	}

	fmt.Println(grid)

	result := 0
	rollsRemoved := -1 // meaning we have not started yet

	for {
		nextGrid := grid
		var remainingRolls int
		for _, row := range nextGrid {
			for j := range row {
				if row[j] == "@" {
					remainingRolls++
				}
			}
		}

		if rollsRemoved == 0 {
			break
		}

		// caputre cordinates of rolls to remove this interation
		// can't simply remove them 1 by 1 cause it will affect the adjacent counts logic
		cordinatesToRemove := [][]int{}

		for i, row := range grid {
			for j := range row {
				if row[j] == "@" {
					adjacentRollCount := 0
					for _, direction := range directions {
						//dir[0] = x
						//dir[1] = y
						nexti := i + direction[0]
						nextj := j + direction[1]

						if nexti < rowSize && nexti >= 0 && nextj < colSize && nextj >= 0 {
							if grid[nexti][nextj] == "@" {
								adjacentRollCount++
							}
						}
					}

					if adjacentRollCount < 4 {
						// figure out all the rolls that are going to be removed in the row:
						cor := make([]int, 2)
						cor[0] = i
						cor[1] = j
						cordinatesToRemove = append(cordinatesToRemove, cor)
						result++
					}
				}

			}
		}

		for _, inx := range cordinatesToRemove {
			nextGrid[inx[0]][inx[1]] = "x"
		}
		rollsRemoved = len(cordinatesToRemove)
	}

	fmt.Printf("====>Result: %d \n", result)

}
