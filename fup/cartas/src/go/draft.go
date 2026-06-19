package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    cartas := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&cartas[i])
    }

    fmt.Print("[")
    for i, c := range cartas {
        if i > 0{
            fmt.Print(", ")
        }
        switch c{
        case 1:
            fmt.Print("A")
        case 11:
            fmt.Print("J")
        case 12:
            fmt.Print("Q")
        case 13:
            fmt.Print("K")
        default:
            fmt.Print(c)
        }
    }
    fmt.Println("]")
}