package main
import "fmt"
func main() {
   var P, D1, D2 int
    fmt.Scan(&P, &D1, &D2)

    soma := D1 + D2
    var vencedor int

    if soma%2 == 0 {
        // soma é par → quem gritou "par" ganha
        vencedor = P
    } else {
        // soma é ímpar → quem gritou "ímpar" ganha
        vencedor = 1 - P
    }
   
   
    fmt.Println(vencedor)
}
