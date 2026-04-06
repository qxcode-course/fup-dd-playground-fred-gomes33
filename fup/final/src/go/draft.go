package main

import "fmt"

func main() {
	var n1, n2, nf, med float64
	fmt.Scan(&n1, &n2, &nf, &med)

	med = (n1 + n2) / 2
	if med >= 7 {
		fmt.Println("aprovado")
	} else if med < 4 {
		fmt.Println("reprovado")
	} else {
		fmt.Println("em recuperação")
	}
}
