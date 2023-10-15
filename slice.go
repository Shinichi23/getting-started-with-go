package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	s := make([]int, 3)
	var input string

	for {
		fmt.Println("Please, add a number, X to quit: ")
		fmt.Scan(&input)
		num, _ := strconv.Atoi(input)

		if input == "X" {
			break
		} else {
			s = append(s, num)
			fmt.Println(s)
			continue
		}

	}
	s = append(s[:0], s[0+3:]...)
	slices.Sort(s)
	fmt.Println(s)
}
