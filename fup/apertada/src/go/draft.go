package main

import "fmt"

func main() {

	array := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}

	menor := array[0]
	for i := 0; i < 5; i++ {
		if array[i] < menor {
			menor = array[i]
		}
	}
	fmt.Println(menor)
}
