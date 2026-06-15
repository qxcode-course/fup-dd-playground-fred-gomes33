package main
import "fmt"
func main() {
    var numE int
    fmt.Scan(&numE)

    Array1 := make([]int, numE)
    for i :=0; i < numE; i++{
        fmt.Scan(&Array1[i])
    }

    var M int
    fmt.Scan(&M)

    Array2 := make([]int, M)
    for i := 0; i < M; i++{
        fmt.Scan(&Array2[i])
    }
    contido := true
    for _, x := range Array1{
        encontrado := false
        for _, y := range Array2{
            if x == y{
                encontrado = true
                break
            }
        }
        if !encontrado{
            contido = false
            break
        }
    }
    if contido{
        fmt.Println("sim")
    }else{
        fmt.Println("nao")
    }
}