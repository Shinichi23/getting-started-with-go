package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Println("Please, enter a floating point number: ")
	fmt.Scan(&number)
	fmt.Println("Truncated integer:", int(number))

}
