package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    N = (N + 1) * (N + 2) / 2
    fmt.Println(N)
}
