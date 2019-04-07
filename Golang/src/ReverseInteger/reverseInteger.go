package main

import (
	"flag"
	"fmt"
)

type flagInt struct {
	value *int32
}

var num int64

func main() {
	fmt.Println("LeetCode Reverse 32 bit Integer problem")
	flag.Int64Var(&num, "num", 0, "Input Integer")
	flag.Parse()

	if num > 2147483647 || num < -2147483648 {
		fmt.Println("Please provide integer within the range [-2147483648, 2147483647]")
		return
	}

	fmt.Println(num)
}

// --------------------------- Solution ---------------------------//

// ------------------------ Helper Function -----------------------//
