package main
import "fmt"
func main() {
    cartela := []int{
        1, 9, 27, 23,
        34, 20, 37, 47,
        30, 87, 55, 69,
        13, 60, 99, 66,
    }

    acertos := 0
    var sorteados int

    for i := 0; i < 6; i++{
        fmt.Scan(&sorteados)

        for _, numero := range cartela{
            if sorteados == numero{
                acertos++
                break
            }
        }
    }
    fmt.Println(acertos)
}