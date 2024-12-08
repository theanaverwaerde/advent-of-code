package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}
	defer f.Close()

	grid := ToGrid(f)

	result := CountXmas(grid)

	fmt.Println("Result:", result)
}

func ToGrid(f *os.File) [][]rune {
	scanner := bufio.NewScanner(f)

	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func CountXmas(grid [][]rune) int {
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != 'X' {
				continue
			}
			count += CountAllDirection(grid, x, y)
		}
	}

	return count
}

func CountAllDirection(grid [][]rune, x, y int) int {
	count := 0

	if IsDirection(grid, x, y, 0, 1) {
		count++
	}
	if IsDirection(grid, x, y, 0, -1) {
		count++
	}
	if IsDirection(grid, x, y, 1, 0) {
		count++
	}
	if IsDirection(grid, x, y, -1, 0) {
		count++
	}
	if IsDirection(grid, x, y, 1, 1) {
		count++
	}
	if IsDirection(grid, x, y, 1, -1) {
		count++
	}
	if IsDirection(grid, x, y, -1, 1) {
		count++
	}
	if IsDirection(grid, x, y, -1, -1) {
		count++
	}

	return count
}

func IsDirection(grid [][]rune, x, y, dirX, dirY int) bool {
	// If OOB
	endX := x + (dirX * 3)
	endY := y + (dirY * 3)
	if endX < 0 || endX >= len(grid[y]) || endY < 0 || endY >= len(grid) {
		return false
	}

	// Detect word (X already tested)
	if grid[y+dirY][x+dirX] != 'M' {
		return false
	}
	if grid[y+(dirY*2)][x+(dirX*2)] != 'A' {
		return false
	}
	if grid[endY][endX] != 'S' {
		return false
	}
	return true
}
