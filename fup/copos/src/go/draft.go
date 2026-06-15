package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    largura := 2*N - 1

    for i := 1; i <=N; i++{
        linha := make([]rune, largura)
        for j := 0; j < largura; j++{
            linha[j] = '.'
        }
        for k := 0; k < i; k++{
            pos := (largura/2) - (i-1) + 2*k
            linha[pos] = rune('0' + N)
        }
        fmt.Println(string(linha))
    }
}