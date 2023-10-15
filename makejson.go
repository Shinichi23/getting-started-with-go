package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please, write your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Please, write your address: ")
	address, _ := reader.ReadString('\n')

	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
	address = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(address, "")

	information := map[string]string{
		"Name":    name,
		"Address": address,
	}

	fmt.Println(information)
	mapToJson, err := json.Marshal(information)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON: ", string(mapToJson))
}
