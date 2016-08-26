package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v, len %d, cap %d", a, len(a), cap(a))

	fmt.Println()

	s1 := a[2:]
	fmt.Printf("%v, len %d, cap %d", s1, len(s1), cap(s1))

	fmt.Println()

	s2 := a[0:8]
	fmt.Printf("%v, len %d, cap %d", s2, len(s2), cap(s2))
}
