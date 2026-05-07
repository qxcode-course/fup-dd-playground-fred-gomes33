package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    if A == B{
        fmt.Println("Empate")
        return
    }

    diff := (B - A + 15) % 15
    if diff >= 1 && diff <= 7{
        fmt.Println("Jogador 1")
        }else {
            fmt.Println("Jogador 2")
        }

}
