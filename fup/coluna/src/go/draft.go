package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    somaColunas := make([]int, n)

    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            var elementos int
            fmt.Scan(&elementos)
            somaColunas[j] += elementos * elementos
        }
    }

    maiorColuna := 0
    maiorSoma := somaColunas[0]
    for j := 1; j < n; j++{
        if somaColunas[j] > maiorSoma{
            maiorSoma = somaColunas[j]
            maiorColuna = j
        }
    }
    fmt.Println(maiorColuna)
}