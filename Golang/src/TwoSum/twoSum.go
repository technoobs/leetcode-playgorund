/*
main package implements the logic to solve LeetCode TwoSum problem
*/

package main

import (
	"fmt"
	"flag"
	"strings"
	"strconv"
)

var numArray string
var result int

func main() {
	fmt.Println("LeetCode Two Sum problem")
	flag.StringVar(&numArray, "num", "", "Number Array")
	flag.IntVar(&result, "result", 0, "Calculation Result")
	flag.Parse()

	fmt.Printf("Number Array is: [%s]\n", numArray)
	fmt.Printf("Calculated Result should be : %d\n", result)

	intArray := generateIntArray(numArray)
	fmt.Printf("asdasd: %d\n", intArray)

	
}

// generateIntArray convert string value into integer array
func generateIntArray(numArray string) []int {
	stringSlice := strings.Split(numArray, " ")
	slice := make([]int, len(stringSlice))
	for index, value := range stringSlice {
		v, _ := strconv.Atoi(value)
		slice[index] = v
	}

	return slice
}