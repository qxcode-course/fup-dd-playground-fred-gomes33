package main
import "fmt"
func main() {
    var rodadas int
    if _, err := fmt.Scan(&rodadas); err != nil || rodadas <= 0{
        return
    }

    valoresReais := make([]float64, rodadas)
    chutesPrimeiro := make([]float64, rodadas)
    escolhasSegundo := make([]string, rodadas)

    for i := 0; i < rodadas; i++{ 
        fmt.Scan(&valoresReais[i])}
    for i := 0; i < rodadas; i++{ 
        fmt.Scan(&chutesPrimeiro[i])}
    for i := 0; i < rodadas; i++{ 
        fmt.Scan(&escolhasSegundo[i])}

    p1, p2 := 0, 0
    for i := 0; i < rodadas; i++{
        vr, cp, es := valoresReais[i], chutesPrimeiro[i], escolhasSegundo[i]

        if cp == vr || (es == "M" && vr < cp) || (es == "m" && vr > cp){
            p1++
        }else {
            p2++
        }
    }
    if p1 > p2{ 
        fmt.Println("primeiro")
    }else if p2 > p1{
        fmt.Println("segundo")
    }else {
        fmt.Println("empate")
    }

}