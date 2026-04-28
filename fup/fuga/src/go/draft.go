package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)
    
    posicao := F

    for {
        if posicao == H{
            fmt.Println("S")
            break
        } else if posicao == P{
            fmt.Println("N")
            break
        }
    posicao = (posicao + D + 16) % 16
    }
}
