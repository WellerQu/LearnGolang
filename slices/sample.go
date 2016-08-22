package main

import "fmt"

func main() {
	s := make([]string, 3)
	//fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//fmt.Println("set: ", s)
	//fmt.Println("get: ", s[2])

	//fmt.Println("len: ", len(s))

	fmt.Println(len(s), cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(len(s), cap(s))

	//fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)

	//l := s[2:5]
	//fmt.Println(len(l), cap(l))
}
