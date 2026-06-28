package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil || n <= 0{
        return
    }

    alturas := make([]float64, n)
    soma := 0.0
    for i := 0; i < n; i++{
        fmt.Scan(&alturas[i])
        soma += alturas[i]
    }

    media := soma / float64(n)
    fmt.Printf("%.2f\n", media)

    for i, alt := range alturas{
        if alt < media{
            fmt.Print("P")
        }else if alt > media {
            fmt.Print("G")
        }else {
            fmt.Print("M")
        }
        if i < n-1{
            fmt.Print(" ")
        }
    }
    fmt.Println()
}