package main
import "fmt"
func main() {
   var n int
   if _, err := fmt.Scan(&n); err != nil {
    return
   }

   tabuleiro := make([][]int, n)
   somaLinha := make([]int, n)

   for i := 0; i < n; i++{
    tabuleiro[i] = make([]int, n)
    for j := 0; j < n; j++{
        fmt.Scan(&tabuleiro[i][j])
        somaLinha[i] += tabuleiro[i][j]
    }
}

pesoMax := 0

for i := 0; i < n; i++{
    for j := 0; j < n; j++{
        somaColunaJ := 0
        for k := 0; k < n; k++{
            somaColunaJ += tabuleiro[k][j]
        }
        pesoAtual := somaLinha[i] + somaColunaJ - (2 * tabuleiro[i][j])
        if pesoAtual > pesoMax{
            pesoMax = pesoAtual
            }
        }
    }
    fmt.Println(pesoMax)
}