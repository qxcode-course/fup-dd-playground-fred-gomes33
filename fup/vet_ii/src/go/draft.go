package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    v := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&v[i])
    }

    fmt.Print("[")
    if n == 0 {
        fmt.Print(" ")
    } else {
        fmt.Print(" ")
        for i := 0; i < n; i++ {
            fmt.Print(v[i])
            if i < n-1 {
                fmt.Print(" ")
            }
        }
        fmt.Print(" ")
    }
    fmt.Println("]")
}