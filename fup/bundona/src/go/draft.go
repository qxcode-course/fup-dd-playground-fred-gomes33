package main
import "fmt"
func main() {
	var h, m int
	var s string
	var d int

	
	_, err := fmt.Scanf("%d\n%d\n%s\n%d\n", &h, &m, &s, &d)
	if err != nil {
		fmt.Scan(&h, &m, &s, &d)
	}

	posInicial := h*6 + m/10

	var posFinal int
	if s == "H" {
		posFinal = (posInicial + d) % 72
	} else {
		posFinal = (posInicial - d) % 72
		if posFinal < 0 {
			posFinal += 72
		}
	}

	horaFinal := posFinal / 6
	minutoFinal := (posFinal % 6) * 10

	fmt.Printf("%02d %02d\n", horaFinal, minutoFinal)
}