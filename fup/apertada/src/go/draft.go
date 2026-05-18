package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var array []int
	array = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}

	menor := array[0]
	for i := 1; i < N; i++ {
		if array[i] < menor {
			menor = array[i]
		}
	}

	fmt.Println(menor)
}
