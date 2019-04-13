package main

import (
	"flag"
	"fmt"
)

func main() {
	var input string
	var result string
	flag.StringVar(&input, "input", "", "Input string value")
	flag.Parse()

	result = solutionOne(input)
	println(result)
}

func solutionOne(input string) string {
	substringStorage := make([]byte, len(input))
	// Map for length and substring
	m := make(map[int][]byte)
	iterate := 0
	for i := range input {
		fmt.Printf("Start character is: %s\n", string(input[i]))
		// Check value "v" exist in the substring storage
		for _, y := range substringStorage {
			if y == input[i] {
				subLength := len(substringStorage)
				fmt.Printf("Find substring without repeated characters: %s\n", string(substringStorage))
				m[iterate] = substringStorage[:subLength]
				substringStorage = substringStorage[:0]
				iterate++
				break
			} else {
				substringStorage[i] = input[i]
				fmt.Printf("Current substring storage slice is: %s\n", string(substringStorage))
			}
		}
	}
	maxLen := 0
	maxIndex := 0
	for x, y := range m {
		if len(y) > maxLen {
			maxLen = len(y)
			maxIndex = x
		}
	}

	return string(m[maxIndex][:])
}
