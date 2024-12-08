package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

		concernRules := GetRulesConcern(values, rules)

		if IsValid(values, concernRules) {
			continue
		}

		reordered := Reorder(values, concernRules)

		result += GetCenterValue(reordered)
	}

	fmt.Println("Result:", result)
}

func GetRule(input string) PageRule {
	values := strings.Split(input, "|")

	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	return PageRule{X: x, Y: y}
}

func Reorder(values []int, rules []PageRule) []int {
	reordered := slices.Clone(values)

	sort.SliceStable(reordered, func(i, j int) bool {
		x := reordered[i]
		y := reordered[j]

		for _, r := range rules {
			if x == r.X && y == r.Y {
				return false
			}

			if x == r.Y && y == r.X {
				return true
			}
		}

		return false
	})

	return reordered
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
		if !RespectRule(values, r) {
			return false
		}
	}

	return true
}

func GetRulesConcern(values []int, rules []PageRule) []PageRule {
	concernRules := []PageRule{}

	for _, r := range rules {
		if IncludingPage(values, r) {
			concernRules = append(concernRules, r)
		}
	}

	return concernRules
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
