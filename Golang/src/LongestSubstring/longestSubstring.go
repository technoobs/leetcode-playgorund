package main

import (
	"flag"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "", "Input string value")
	flag.Parse()
}
