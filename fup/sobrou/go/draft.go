package main
import "fmt"
func main() {
   var qu1, qu2, qu3, val1, val2, val3, money float32
   fmt.Scan(&qu1, &qu2, &qu3, &val1, &val2, &val3, &money)
   var S float32 = (
    (money - ((qu1 * val1) + (qu2 * val2) + (qu3 * val3)))) 
   fmt.Printf("%.2f\n", S)
}
