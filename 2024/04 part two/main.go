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

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] != 'A' {
				continue
			}
			count += CountAllDirection(grid, x, y)
		}
	}

	return count
}

func CountAllDirection(grid [][]rune, x, y int) int {
	count := 0

	if ShapeOne(grid, x, y) {
		count++
	}
	if ShapeTwo(grid, x, y) {
		count++
	}
	if ShapeThree(grid, x, y) {
		count++
	}
	if ShapeFour(grid, x, y) {
		count++
	}

	return count
}

func ShapeOne(grid [][]rune, x, y int) bool {
	// M•M
	// •A•
	// S•S

	if grid[y+1][x+1] != 'M' {
		return false
	}
	if grid[y+1][x-1] != 'M' {
		return false
	}
	if grid[y-1][x+1] != 'S' {
		return false
	}
	if grid[y-1][x-1] != 'S' {
		return false
	}
	return true
}

func ShapeTwo(grid [][]rune, x, y int) bool {
	// M•S
	// •A•
	// M•S

	if grid[y+1][x-1] != 'M' {
		return false
	}
	if grid[y-1][x-1] != 'M' {
		return false
	}
	if grid[y+1][x+1] != 'S' {
		return false
	}
	if grid[y-1][x+1] != 'S' {
		return false
	}
	return true
}

func ShapeThree(grid [][]rune, x, y int) bool {
	// S•S
	// •A•
	// M•M

	if grid[y-1][x+1] != 'M' {
		return false
	}
	if grid[y-1][x-1] != 'M' {
		return false
	}
	if grid[y+1][x+1] != 'S' {
		return false
	}
	if grid[y+1][x-1] != 'S' {
		return false
	}
	return true
}

func ShapeFour(grid [][]rune, x, y int) bool {
	// S•M
	// •A•
	// S•M

	if grid[y+1][x+1] != 'M' {
		return false
	}
	if grid[y-1][x+1] != 'M' {
		return false
	}
	if grid[y+1][x-1] != 'S' {
		return false
	}
	if grid[y-1][x-1] != 'S' {
		return false
	}
	return true
}
