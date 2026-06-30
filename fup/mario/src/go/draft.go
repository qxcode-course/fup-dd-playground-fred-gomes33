package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil || n < 0 {
        return
    }
    alturas := make([]int, n)
    maxAltura := 0

    for i := 0; i < n; i++ {
        if _, err := fmt.Scan(&alturas[i]); err != nil {
            return
        }
        if alturas[i] > maxAltura {
            maxAltura = alturas[i]
        }
    }

    for linha := maxAltura; linha >= 1; linha--{
        for coluna := 0; coluna < n; coluna++{
            if alturas[coluna] >= linha {
                fmt.Print("#")
            }else{
                fmt.Print("_")
            }
        }
        fmt.Println()
    }
}