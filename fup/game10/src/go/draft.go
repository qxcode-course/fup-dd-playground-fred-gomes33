package main
import "fmt"
func main() {
   var N, D, A int
   fmt.Scan(&N, &D, &A)

   //Cálculo num. min.
   mov := (D - A + N) % N

   fmt.Println(mov)
}
