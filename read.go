package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	fmt.Println("Hello, please write your filename")

	inputReader := bufio.NewReader(os.Stdin)
	fileName, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fileName = strings.ReplaceAll(fileName, "\n", "")

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	names := make([]name, 0, 3)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var aName name
		aName.fname, aName.lname = s[0], s[1]
		names = append(names, aName)

	}

	i := 0

	for range names {
		fmt.Println(names[i])
		i++
	}

}
