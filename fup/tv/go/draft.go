package main
import "fmt"
func main() {
    var valtv, quap, valp, valt float64
    fmt.Scan(&valtv, &quap, &valp, &valt)
    valp = valtv * quap
    valt = valp * 0.15  
    

    
fmt.Printf("%.2f\n%.2f\n", valp, valt)

}