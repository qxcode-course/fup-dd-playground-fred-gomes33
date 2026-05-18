package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    Array := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&Array[i])
    }
    if N == 0{
        fmt.Println("[]")
        return
    }
    fmt.Print("[")
    
    for i := 0; i < N; i++{
        fmt.Print(Array[i])
        if i < N-1{
            fmt.Print(", ")
        }
    }
    fmt.Println("]")
}
