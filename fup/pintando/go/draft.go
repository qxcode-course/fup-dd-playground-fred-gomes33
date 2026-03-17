package main
import "fmt"
import "math"
func main() {
    var l1, l2, l3, tt, AT, raiz float64
    fmt.Scan(&l1, &l2, &l3, &tt, &AT, &raiz)
   tt = (l1 + l2 + l3) / 2
   AT = (tt * ((tt - l1) * (tt - l2) * (tt - l3)))
   raiz = math.Sqrt(AT)
   fmt.Printf("%.2f\n", raiz)
}
