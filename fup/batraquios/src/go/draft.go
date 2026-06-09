package main
import "fmt"
func main() {
    var numE int
    fmt.Scan(&numE)

    Array := make([]int, numE)
    Array2 := make([]int, numE)
    for i :=0; i < numE; i++{
        fmt.Scan(&Array[i], &Array2[i])
        if Array[i] == Array2[i] {
            fmt.Print("sim")
        }else{
            fmt.Print("nao")
        }
    }
}