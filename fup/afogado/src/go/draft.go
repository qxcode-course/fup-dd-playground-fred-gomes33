package main
import "fmt"
func main() {
    var P, E int
    fmt.Scan(&P, &E)

    for forçaInicial := 1; ; forçaInicial++ {
        altura := 0
        salto := forçaInicial

        for {
            altura += salto
            if altura >= P {
                fmt.Println(forçaInicial)
                return
            }
            altura -= E
            if altura < 0 {
                break
            }
            salto -= 10
            if salto <= 0 {
                break
            }
        }
    }
}
