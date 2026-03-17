package main
import "fmt"
import "math"
func main() {
  var x1, y1, x2, y2, d float64
  fmt.Scan(&x1, &y1, &x2, &y2, &d)  
  d = math.Hypot(x2-x1, y2-y1)
 d = math.Round(d*100)/100
  fmt.Printf("%.2f\n", d)

}
