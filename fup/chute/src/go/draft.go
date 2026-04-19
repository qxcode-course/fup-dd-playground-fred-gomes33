package main
import "fmt"
import "math"
func main() {
   var val, pc, sc int
   fmt.Scan(&val, &pc, &sc)

   distprim := int(math.Abs(float64(val - pc)))
   distsec := int(math.Abs(float64(val - sc)))

    if distprim < distsec {
         fmt.Println("primeiro")
    } else if distsec < distprim {
         fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }
}
