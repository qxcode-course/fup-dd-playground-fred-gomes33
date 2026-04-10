package main
import "fmt"
func main() {
var B, T int
    fmt.Scan(&B, &T)

    
    larg := 160
    alt := 70

    
    aT := larg * alt

    // Área do pedaço do Felix (trapézio da esquerda)
    aFelix := (B + T) * alt / 2

    // Área do pedaço da Marzia
    aMarzia := aT - aFelix

    // Metade da área total
    met := aT / 2

    // Decisão
    if aFelix > met {
        fmt.Println(1) // Felix ficou com o pedaço que vale 100 reais
    } else if aMarzia > met {
        fmt.Println(2) // Marzia ficou com o pedaço que vale 100 reais
    } else {
        fmt.Println(0) // O valor da nota se perdeu
    }

}
