package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	Obstacle  bool
	Parcoured bool
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}
	defer f.Close()

	grid, startX, startY := ToGrid(f)

	Patrol(grid, startX, startY)

	count := CountDistinctPosition(grid)

	fmt.Println("Result:", count)
}

func ToGrid(f *os.File) ([][]Tile, int, int) {
	scanner := bufio.NewScanner(f)

	var grid [][]Tile

	var startX, startY int

	y := 0
	for scanner.Scan() {
		grid = append(grid, []Tile{})

		for x, value := range scanner.Text() {
			if value == '^' {
				startX = x
				startY = len(grid) - 1
			}

			grid[y] = append(grid[y], Tile{Obstacle: value == '#'})
		}
		y++
	}

	return grid, startX, startY
}

func CountDistinctPosition(grid [][]Tile) int {
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].Parcoured {
				count++
			}
		}
	}
	return count
}

func Patrol(grid [][]Tile, x, y int) {
	// When panic patrol is finish
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Patrol finish")
		}
	}()

	direction := 'N'

	for {
		grid[y][x].Parcoured = true

		dirX, dirY := GetDirection(direction)

		if grid[y+dirY][x+dirX].Obstacle {
			direction = ChangeDirection(direction)
		}

		dirX, dirY = GetDirection(direction)

		x += dirX
		y += dirY
	}
}

func GetDirection(dir rune) (int, int) {
	switch dir {
	case 'N':
		return 0, -1
	case 'S':
		return 0, 1
	case 'E':
		return 1, 0
	case 'W':
		return -1, 0
	}
	panic("Not valid direction")
}

func ChangeDirection(dir rune) rune {
	switch dir {
	case 'N':
		return 'E'
	case 'E':
		return 'S'
	case 'S':
		return 'W'
	case 'W':
		return 'N'
	}
	panic("Not valid direction")
}
