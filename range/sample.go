package main

import "fmt"

func main() {
	fmt.Println("Hello Range")

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}

	kvs := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	for k, v := range kvs {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
