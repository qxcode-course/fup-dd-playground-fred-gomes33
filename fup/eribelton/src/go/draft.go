package main
import "fmt"
func obterValorAscologico(palavra string) int {
	soma := 0
	for _, c := range palavra{
        soma += int(c-'a')
    }
	return soma % 2
}

func main() {
	var palavra string
	fmt.Scan(&palavra)

	valorOriginal := obterValorAscologico(palavra)
	menorValor := valorOriginal
	melhorPalavra := palavra
	for c := 'a'; c <= 'z'; c++ {
		nova := palavra + string(c)
        v := obterValorAscologico(nova)

		if v < menorValor {
			menorValor = v
			melhorPalavra = nova
		}
	}

	fmt.Println(valorOriginal)
	fmt.Println(melhorPalavra)
	fmt.Println(menorValor)
}