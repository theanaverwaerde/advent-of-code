package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PageRule struct {
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

	resolving := false

	rules := []PageRule{}

	result := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			resolving = true
			continue
		}

		if !resolving {
			rules = append(rules, GetRule(input))
			continue
		}

		values := GetValues(input)

		if IsValid(values, rules) {
			result += GetCenterValue(values)
		}
	}

	fmt.Println("Result:", result)
}

func GetRule(input string) PageRule {
	values := strings.Split(input, "|")

	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	return PageRule{X: x, Y: y}
}

func GetValues(input string) []int {
	rawValues := strings.Split(input, ",")

	values := make([]int, len(rawValues))

	for i, v := range rawValues {
		x, _ := strconv.Atoi(v)

		values[i] = x
	}

	return values
}

func GetCenterValue(values []int) int {
	center := len(values) / 2

	return values[center]
}

func IsValid(values []int, rules []PageRule) bool {

	for _, r := range rules {
		if !IncludingPage(values, r) {
			continue
		}

		if !RespectRule(values, r) {
			return false
		}
	}

	return true
}

func IncludingPage(values []int, rule PageRule) bool {
	var haveX, haveY bool

	for _, v := range values {
		if v == rule.X {
			haveX = true
		}
		if v == rule.Y {
			haveY = true
		}

		if haveX && haveY {
			return true
		}
	}

	return false
}

func RespectRule(values []int, rule PageRule) bool {
	var pass bool

	for _, v := range values {
		if v == rule.Y {
			return pass
		}

		if v == rule.X {
			pass = true
		}
	}

	panic("not found rule Y")
}
