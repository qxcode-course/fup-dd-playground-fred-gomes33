package main

import "fmt"

func main() {
	    /*var lista []int = []int{1, 2, 3, 4, 5} // lista estática
		var nomes []string = []string{"Alice", "Bob", "Charlie"}
		nomes = []string{"Emily", "Clara", "Biel"}
		fmt.Println(nomes)
	    fmt.Println(nomes[0]) // imprime "Emily" que está no índice/posição 0
		fmt.Println(lista[2]) // imprime "3" que está no índice/posição 2*/

        arr := []int{9, 8, 4, 5, 6, 2, 3} // de trás para frente
        fmt.Print("[ ")
        for i := len(arr) - 1; i >= 0; i-- {
        fmt.Printf("%d ", arr[i])
        }
        fmt.Print(" ]\n")


        arra := []int{9, 8, 4, 5, 6, 2, 3} // de trás para frente
        fmt.Print("[ ")
        for i := len(arra) - 1; i >= 0; i-- {
        fmt.Printf("%d ", arra[i])
        }
        fmt.Print(" ]\n")

        array := []int{9, 8, 4, 5, 6, 2, 3} 
        fmt.Print("[ ")
        for _, valor := range array{
            fmt.Printf("%d ", valor)
        }
        fmt.Print(" ]\n")
    }
