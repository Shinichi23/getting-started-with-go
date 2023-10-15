package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
		G
		H
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
}
