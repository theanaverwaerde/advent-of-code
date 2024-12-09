package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("No input.txt file")
		return
	}

	idNumber := GetIdNumber(string(content))

	fmt.Println("ID number:", idNumber)

	rearranged := Arranged(idNumber)

	fmt.Println("Arranged value:", rearranged)

	checksum := GetChecksum(rearranged)

	fmt.Println("Checksum:", checksum)
}

func GetIdNumber(input string) []int {
	idNumber := []int{}

	for i,v := range input {
		vInt, _ := strconv.Atoi(string(v))

		for j := 0; j < vInt; j++ {
			if i%2 == 0 {
				idNumber = append(idNumber, i/2)
			} else {
				idNumber = append(idNumber, -1)
			}
		}
	}

	return idNumber
}

func Arranged(input []int) []int {
	output := []int{}

	start := 0
	end := len(input)-1

	for ; start <= end; {
		if input[end] == -1 {
			end--
			continue
		}

		if input[start] == -1 {
			output = append(output, input[end])
			end--
		} else {
			output = append(output, input[start])
		}

		start++
	}
	

	return output
}

func GetChecksum(input []int) int {
	result := 0

	for i,v := range input {
		result += i * v
	}

	return result
}