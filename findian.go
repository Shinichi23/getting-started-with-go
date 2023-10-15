package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please, enter a string: ")
	str, _ := reader.ReadString('\n')

	lowerCase := strings.ToLower(strings.TrimSpace(str))

	start := strings.HasPrefix(lowerCase, "i")
	contain := strings.Contains(lowerCase, "a")
	end := strings.HasSuffix(lowerCase, "n")

	if start && end && contain {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
