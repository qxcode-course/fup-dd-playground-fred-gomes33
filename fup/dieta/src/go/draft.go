package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    Array := make([]int, N)
    soma := 0
    for i := 0; i < N; i++{
        fmt.Scan(&Array[i])
        soma += Array[i]
    }

med := float64(soma) / float64(N)
    fmt.Printf("%.1f\n", med)
}
