package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    esq := make([]int, 70)
    dir := make([]int, 70)

    for i := 0; i < n; i++ {

        var tamanho int
        var lado string

        fmt.Scan(&tamanho, &lado)

        if lado == "E" {
            esq[tamanho]++
        } else {
            dir[tamanho]++
        }
    }

    pares := 0

    for i := 0; i < 70; i++ {

        if esq[i] < dir[i] {
            pares += esq[i]
        } else {
            pares += dir[i]
        }
    }

    fmt.Println(pares)
}