package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var xList, yList []int

	for scanner.Scan() {
		x, y := GetValues(scanner.Text())
		xList = append(xList, x)
		yList = append(yList, y)
	}

	slices.Sort(xList)
	slices.Sort(yList)

	var result int

	for i := 0; i < len(xList); i++ {
		result += GetDistance(xList[i], yList[i])
	}

	fmt.Println("Result:", result)
}

func GetDistance(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func GetValues(s string) (int, int) {
	values := strings.Split(s, "   ")

	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	return x, y
}
