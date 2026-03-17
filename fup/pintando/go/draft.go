package main
import "fmt"
func main() {
    var l1, l2, l3, tt float64
    fmt.Scan(&l1, &l2, &l3, &tt)
   tt = (l1 + l2 + l3) / 2
    fmt.Printf("%.2f\n", tt)
}
