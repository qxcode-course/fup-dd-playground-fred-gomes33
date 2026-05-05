package main

import "fmt"

func mostrar_vetor(arr []int, sep string) {
	for i, valor := range arr {
		if i != 0 {
			fmt.Println(sep)
		}
		fmt.Printf("%v", valor)
	}
	
}

func main() {
	var qtd int
	fmt.Scan(&qtd)
	var arr []int = make([]int, qtd)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	mostrar_vetor(arr, " ")
}
