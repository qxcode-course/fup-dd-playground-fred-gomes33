package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil{
        return
    }

    vetor := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&vetor[i])
    }

    for i := 0; i < n/2; i++{
        vetor[i], vetor[n-1-i] = vetor[n-1-i], vetor[i]
    }

    fmt.Print("[")
    for _, num := range vetor{
        fmt.Printf(" %d", num)
    }
    fmt.Println(" ]")
}