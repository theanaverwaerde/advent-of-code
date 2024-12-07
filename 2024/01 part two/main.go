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

	var xList, yList []int

	for scanner.Scan() {
		x, y := GetValues(scanner.Text())
		xList = append(xList, x)
		yList = append(yList, y)
	}

	var result int

	for _, x := range xList {
		var occurence int

		occurence += GetOccurence(x, yList)

		result += x * occurence
	}

	fmt.Println("Result:", result)
}

func GetOccurence(x int, yList []int) int {
	count := 0

	for _, y := range yList {
		if x == y {
			count++
		}
	}

	return count
}

func GetValues(s string) (int, int) {
	values := strings.Split(s, "   ")

	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	return x, y
}
