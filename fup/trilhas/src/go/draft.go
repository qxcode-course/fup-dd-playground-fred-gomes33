package main
import "fmt"
import "math"
func main() {
   var n int
   if _, err := fmt.Scan(&n); err != nil {
     return
   }

   melhorTrilha := 1
   menorEsforco := math.MaxInt32

   for t := 1; t <= n; t++{
    var m int
    fmt.Scan(&m)

    h := make([]int, m)
    for i := 0; i < m; i++{
        fmt.Scan(&h[i])
    }

    subidaIda := 0
    for i := 0; i < m-1; i++{
      if h[i+1] > h[i]{
        subidaIda += h[i+1] - h[i]
      }
    }
    subidaVolta := 0
    for i := m - 1; i > 0; i--{
      if h[i-1] > h[i]{
        subidaVolta += h[i-1] - h[i]
      }
    }
    esforcoTrilha := subidaIda
    if subidaVolta < subidaIda {
      esforcoTrilha = subidaVolta
    }

    if esforcoTrilha < menorEsforco{
      menorEsforco = esforcoTrilha
      melhorTrilha = t
    }
   }
    fmt.Println(melhorTrilha)
}