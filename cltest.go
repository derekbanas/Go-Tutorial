package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	// Get all values after the first index
	args := os.Args[1:]

	// Create int array from string array
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value :", max)
}
