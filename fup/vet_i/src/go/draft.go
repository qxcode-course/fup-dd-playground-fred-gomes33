package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n) // -> criar um array inteiro de tamanho n(como se fosse o x/y na mat)
	
	if n == 0 {
		fmt.Println("")
	}
	for i := 0; i < n; i++{ // lê os elementos e guarda no array
		fmt.Scan(&array[i])
	}

	for _, v := range array{// imprime elemntos na ordem de leitura
		fmt.Println(v)      // range -> devolve indice e valor
	}
}