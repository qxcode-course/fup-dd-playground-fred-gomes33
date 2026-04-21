package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Printf("[ ")
	if A < B {
		//crescente
        for i := A; i < B; i++ {
			fmt.Printf("%d ", i)
		}
	} else {
		//decrescente
		for i := A; i > B; i-- {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("]\n")
}
