package main
import "fmt"
func main() {
   
   var N, X, Y, S int
   var C string
   fmt.Scan(&N, &X, &Y, &C, &S)

   switch C {
    case "R":
        X = (X + S) % N
    case "L":
        X = (X - S) % N
        if X < 0 {
            X += N
        }
    case "D":
        Y = (Y + S) % N
    case "U":
        Y = (Y - S) % N
        if Y < 0 {
            Y += N
        }
   }  
   
    fmt.Printf("%d %d\n", X, Y)
}
