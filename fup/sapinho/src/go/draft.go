package main
import "fmt"
func main() {
    var P, S, E int
    fmt.Scan(&P, &S, &E)
    //P -> Profundidade do poço 
    //S -> Altura do salto
    //E -> Queda
    // Sendo sempre S>E

    posicao := 0

    for {
        salto := posicao + S
        if salto >= P{
            fmt.Printf("%d saiu\n", posicao)
            break
        }
        fmt.Printf("%d %d\n", posicao, salto)

        posicao = salto - E

    }
}
