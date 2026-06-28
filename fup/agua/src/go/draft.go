package main
import "fmt"
func main() {
    var n, q int
    if _, err := fmt.Scan(&n, &q); err != nil{
        return
    }

    casa := make([]int, n)

    for i := 0; i < q; i++{
        var a, b, l int
        fmt.Scan(&a, &b, &l)

        for j := a; j <= b; j++{
            casa[j] += l
        }
}

for i, v := range casa{
    fmt.Print(v)
    if i < n-1{
        fmt.Print(" ")
    }
}
    fmt.Println()
}