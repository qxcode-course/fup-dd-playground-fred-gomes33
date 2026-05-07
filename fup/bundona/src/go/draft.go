package main
import "fmt"
func main() {
  var H, M int
    var S string
    var D int

    // Entrada
    fmt.Scan(&H)
    fmt.Scan(&M)
    fmt.Scan(&S)
    fmt.Scan(&D)

    index := H*6 + M/10

    passos := D

    if S == "H" {
        index = (index + passos) % 72
    } else {
        index = (index - passos + 72) % 72
    }

    hora := index / 6
    minuto := (index % 6) * 10

    fmt.Printf("%02d %02d\n", hora, minuto)
}
