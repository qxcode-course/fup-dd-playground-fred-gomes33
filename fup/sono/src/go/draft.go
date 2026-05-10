package main
import "fmt"
func main() {
    var h1, m1, s1 int // hora atual
    var h2, m2, s2 int // hora de acordar

    fmt.Scanf("%d %d %d", &h1, &m1, &s1)
    fmt.Scanf("%d %d %d", &h2, &m2, &s2)

    dorminS := h1*3600 + m1*60 + s1
    acordarS := h2*3600 + m2*60 + s2

    if acordarS < dorminS {
        acordarS += 24 * 3600
    }

    duracaoS := acordarS - dorminS

    h := duracaoS / 3600
    m := (duracaoS % 3600) / 60
    s := duracaoS % 60
    fmt.Printf("%02d %02d %02d\n", h, m, s)
}

