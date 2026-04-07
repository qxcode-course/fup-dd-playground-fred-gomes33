package main

import "fmt"

func main() {
	var n1, n2, nf, med, medfd int
	fmt.Scan(&n1, &n2, &nf, &med, &medfd)

	med = (n1 + n2) / 2
	if med >= 7 {
		fmt.Println("aprovado")
	} else if med < 4 {
		fmt.Println("reprovado")
	} else if med >= 4 && med < 7 {
		medfd = (med + nf) / 2
		if medfd >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	} else {
		fmt.Println("reprovado na final")
	}

}
