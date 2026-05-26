package main
import "fmt"
type Bota struct {
	tam int
	pe  string
}

func gerarPar(bota Bota) Bota {
	par := Bota{tam: bota.tam}
	if bota.pe == "D" {
		par.pe = "E"
	} else {
		par.pe = "D"
	}
	return par
}

func procurarBota(lista []Bota, bota Bota) (bool, int) {
	for i, elem := range lista {
		if elem == bota {
			return true, i
		}
	}
	return false, 0
}

func main() {
	qtd := 0
	fmt.Scan(&qtd)
	botas := make([]Bota, qtd)
	for i := range botas {
		fmt.Scan(&botas[i].tam, &botas[i].pe)
	}
	descasados := make([]Bota, 0)
	cont_pares := 0
	for _, elem := range botas {
		par := gerarPar(elem)
		encontrei, onde := procurarBota(descasados, par)
		if !encontrei {
			descasados = append(descasados, elem)
		} else {
			cont_pares += 1
			descasados[onde].tam = 0
			// slices.Delete(descasados, onde, onde + 1)
		}
	}
	fmt.Println(cont_pares)
}
 