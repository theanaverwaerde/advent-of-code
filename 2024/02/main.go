package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	safeCount := 0

	for scanner.Scan() {
		if IsSafe(scanner.Text()) {
			safeCount++
		}
	}

	fmt.Println("Safe reports count:", safeCount)
}

func IsSafe(s string) bool {
	values := strings.Split(s, " ")

	way := DetectWay(values)

	if way == 0 {
		return false
	}

	for i := 1; i < len(values); i++ {
		x, _ := strconv.Atoi(values[i-1])
		y, _ := strconv.Atoi(values[i])

		distance := (x - y) * way

		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true
}

func DetectWay(values []string) int {
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	number := x - y

	if number < 0 {
		return -1
	} else if number > 0 {
		return 1
	} else {
		return 0
	}
}
