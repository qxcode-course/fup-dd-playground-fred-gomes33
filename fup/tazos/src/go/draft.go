package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    tazos := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&tazos[i])
    }

    maxCount := 1
    count := 1
    resultado := []int{tazos[0]}

    for i := 1; i < n; i++{
        if tazos[i] == tazos[i-1]{
            count++
        }else {
            count = 1
        }
        if count > maxCount{
            maxCount = count
            resultado = []int{tazos[i]}
        }else if count == maxCount{
            if resultado[len(resultado)-1] != tazos[i]{
                resultado = append(resultado, tazos[i])
            }
        }
    }
    
    fmt.Print("[ ")
    for i, v := range resultado{
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println(" ]")
}