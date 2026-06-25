package main
import "fmt"
func main() {
	var totalAlbum, totalPossuidas int
	
	if _, err := fmt.Scan(&totalAlbum); err != nil {
		return
	}
	if _, err := fmt.Scan(&totalPossuidas); err != nil {
		return
	}

	frequencia := make([]int, totalAlbum+1)
	var repetidas []int

	for i := 0; i < totalPossuidas; i++ {
		var fig int
		fmt.Scan(&fig)

		if frequencia[fig] > 0 {
			repetidas = append(repetidas, fig)
		}
		frequencia[fig]++
	}

	// Identificação das figurinhas que estão faltando
	var faltando []int
	for i := 1; i <= totalAlbum; i++ {
		if frequencia[i] == 0 {
			faltando = append(faltando, i)
		}
	}

	imprimirResultado(repetidas)
	imprimirResultado(faltando)
}

func imprimirResultado(lista []int) {
	fmt.Print("[")
	for _, num := range lista {
		fmt.Printf(" %d", num)
	}
	
	fmt.Println(" ]")
}