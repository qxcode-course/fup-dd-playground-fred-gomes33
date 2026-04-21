package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    fmt.Printf("[ ")
    for i := A; i < B; i++{
        fmt.Printf("%d ", i)
    }
    fmt.Printf("]\n")
}
