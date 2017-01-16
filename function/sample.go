package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func main() {

	fmt.Println(plus(1, 2))
	fmt.Println(minus(2, 1))
}
