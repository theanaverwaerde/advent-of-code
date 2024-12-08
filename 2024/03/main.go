package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}

	result := Process(string(content))

	fmt.Println("Result:", result)
}

func Process(input string) int {
	var re = regexp.MustCompile(`(?m)mul\((\d{1,3}),(\d{1,3})\)`)

	result := 0

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		result += x * y
	}
	return result
}
