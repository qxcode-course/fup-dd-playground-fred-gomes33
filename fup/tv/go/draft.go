package main
import "fmt"
func main() {
    var valor float64
    var parcelas int

    fmt.Scanln(&valor)
    fmt.Scanln(&parcelas)

    var taxa float64
    switch parcelas {
    case 1:
        taxa = 0.0
    case 2:
        taxa = 0.05
    case 3:
        taxa = 0.10
    case 4:
        taxa = 0.15
    case 5:
        taxa = 0.20
    case 6:
        taxa = 0.25
    case 7:
        taxa = 0.30
    case 8:
        taxa = 0.35
    case 9:
        taxa = 0.40
    case 10:
        taxa = 0.45

    }
valorTotal := valor * (1 + taxa)
    valorParcela := valorTotal / float64(parcelas)

    fmt.Printf("%.2f\n", valorParcela)
    fmt.Printf("%.2f\n", valorTotal)
}