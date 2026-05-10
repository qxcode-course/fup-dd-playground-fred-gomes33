package main
import "fmt"
import "sort"

func main() {
    var np1, np2, np3, ntr float64
    fmt.Scan(&np1, &np2, &np3, &ntr)

    provas := []float64{np1, np2, np3}
    sort.Float64s(provas)

    somar := provas[1] + provas[2]
    media := (somar + ntr) / 3.0

    if media >= 7.0{
        fmt.Printf("Aprovado com %.1f\n", media)
    } else {
        fmt.Printf("Final com %.1f\n", media)

    }
 
}
