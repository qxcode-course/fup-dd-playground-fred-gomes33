package main
import "fmt"
func main() {
    var l, n int
	if _, err := fmt.Scan(&l, &n); err != nil {
		return
	}

	ganhador, perdedor := -1, 0
	var menorDist, maiorDist int

	for i := 0; i < n; i++ {
		var jogada int
		fmt.Scan(&jogada)

		dist := jogada
		if dist < 0 {
			dist = -dist
		}

		
		if dist <= l {
			if ganhador == -1 || dist <= menorDist {
				menorDist = dist
				ganhador = i
			}
		}

		
		if i == 0 || dist >= maiorDist {
			maiorDist = dist
			perdedor = i
		}
	}

	if ganhador == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(ganhador)
	}
	fmt.Println(perdedor)
}