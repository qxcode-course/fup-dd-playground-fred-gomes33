package main
import "fmt"
func main() {
    var nElementos int
    fmt.Scan(&nElementos)

    Array := make([]int, nElementos)

    for i := 0; i < nElementos; i++{
        fmt.Scan(&Array[i])
    }

    forcaJ := 0
    forcaS := 0

    for i := 0; i < nElementos; i++{
        if i < (nElementos/2){
            forcaJ += Array[i]
        }else if i >= (nElementos/2){
            forcaS += Array[i]
        }
    }

    if forcaJ > forcaS {
        fmt.Println("Jedi")
    }else if forcaS > forcaJ{
        fmt.Println("Sith")
    }else if forcaS == forcaJ{
        fmt.Println("Empate")
    }
}