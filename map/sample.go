package main

import "fmt"

func main() {
	fmt.Println("Hello Map!")

	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 4

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	_, second := m["k1"]
	fmt.Println("second:", second)

	_, nothere := m["k3"]
	fmt.Println("nothere:", nothere)

	m2 := map[string]int{"key1": 123, "key2": 456}
	fmt.Println(m2)
}
