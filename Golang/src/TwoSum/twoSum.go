/*
main package implements the logic to solve LeetCode TwoSum problem
*/

package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var numArray string
var result int
var solution int

func main() {
	fmt.Println("LeetCode Two Sum problem")
	flag.StringVar(&numArray, "num", "", "Number Array")
	flag.IntVar(&result, "result", 0, "Calculation Result")
	flag.IntVar(&solution, "solution", 0, "Solution")
	flag.Parse()

	intArray := generateIntArray(numArray)

	// logic
	switch solution {
	case 1:
		fmt.Println("Solution 1")
		fmt.Printf("Return %d", solutionOne(intArray, result))
	case 2:
		fmt.Println("Solution 2")
		fmt.Printf("Return %d", solutionTwo(intArray, result))
	default:
		fmt.Println("Solution not imeplemented")
	}
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

// Solution 01
func solutionOne(numbers []int, expected int) []int {
	result := make([]int, 2)
	for i1, v1 := range numbers {
		s := numbers[i1+1:]
		for i2, v2 := range s {
			if v1 >= expected || v2 >= expected {
				continue
			}
			if v1+v2 == expected {
				result[0] = i1
				result[1] = i1 + 1 + i2
				break
			}
		}
	}
	return result
}

// Solution 02
func solutionTwo(numbers []int, expected int) []int {
	result := make([]int, 2)
	for i1, v1 := range numbers {
		if v1 > expected {
			continue
		}
		a := expected - v1
		s := numbers[i1+1:]
		for i2, v2 := range s {
			if a != v2 {
				continue
			} else {
				result[0] = i1
				result[1] = i1 + 1 + i2
			}
		}
	}
	return result
}
