package main

import "fmt"

func main() {
	var P, S, E int
	if _, err := fmt.Scan(&P, &S, &E); err != nil {
		return
	}

	posicao := 0
	salto := S

	for {
		posSalto := posicao + salto
		if salto < 0 {
			posSalto = posicao + salto + 10
		}

		if posSalto >= P {
			fmt.Printf("%d saiu\n", posicao)
			break
		}

		fmt.Printf("%d %d\n", posicao, posSalto)

		posicao = posSalto - E

		if posicao < 0 {
			fmt.Printf("%d morreu\n", posicao)
			break
		}
		salto -= 10
	}
}