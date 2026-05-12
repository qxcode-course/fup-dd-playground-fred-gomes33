package main
import "fmt"

func main() {
    var P, S, E int
    fmt.Scan(&P, &S, &E)

    posicao := 0
    salto := S

    for {
        posSalto := posicao + salto

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
        if salto <= 0 {
            posicao = posicao - E
            fmt.Printf("%d morreu\n", posicao)
            break
        }
    }
}
