package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}
	defer f.Close()

	grid := ToGrid(f)

	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != 0 {
				continue
			}

			nineCount := CountToNine(grid, x, y, 1)

			noDuplicate := RemoveDuplicatePosition(nineCount)

			count += len(noDuplicate)
		}
	}

	fmt.Println("Sum:", count)
}

func ToGrid(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)

	var grid [][]int

	for scanner.Scan() {

		grid = append(grid, ToIntSlice(scanner.Text()))
	}

	return grid
}

func ToIntSlice(input string) []int {
	output := []int{}

	for _, v := range input {
		vInt, _ := strconv.Atoi(string(v))
		output = append(output, vInt)
	}

	return output
}

func CountToNine(grid [][]int, x, y int, search int) []Position {
	if GetValue(grid, x, y) == 9 {
		return []Position{{X: x, Y: y}}
	}

	nineCount := []Position{}

	nextSearch := search + 1

	// Top
	if GetValue(grid, x, y-1) == search {
		nineCount = append(nineCount, CountToNine(grid, x, y-1, nextSearch)...)
	}

	// Bottom
	if GetValue(grid, x, y+1) == search {
		nineCount = append(nineCount, CountToNine(grid, x, y+1, nextSearch)...)
	}

	// Right
	if GetValue(grid, x+1, y) == search {
		nineCount = append(nineCount, CountToNine(grid, x+1, y, nextSearch)...)
	}

	// Left
	if GetValue(grid, x-1, y) == search {
		nineCount = append(nineCount, CountToNine(grid, x-1, y, nextSearch)...)
	}

	return nineCount
}

func GetValue(grid [][]int, x, y int) int {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return -1
	}

	return grid[y][x]
}

func RemoveDuplicatePosition(positions []Position) []Position {
	uniquePositions := make(map[Position]bool)
	var result []Position

	for _, pos := range positions {
		if !uniquePositions[pos] {
			uniquePositions[pos] = true
			result = append(result, pos)
		}
	}

	return result
}
