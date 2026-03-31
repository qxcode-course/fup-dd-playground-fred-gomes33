package main
import "fmt"
func main() {
    var vm, tempmin, cons, temph, dist, desem float64
    fmt.Scan(&vm, &tempmin, &cons, &temph, &dist, &desem)
    temph = tempmin / 60
    dist = vm * temph
    desem = dist / cons

    fmt.Printf("%.2f\n", desem)
}
