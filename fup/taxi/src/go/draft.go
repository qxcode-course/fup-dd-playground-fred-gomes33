package main
import "fmt"
func main() {
   var A, G, Ra, Rg float64
   fmt.Scanf("%f\n%f\n%f\n%f", &A, &G, &Ra, &Rg)

   cAlcool := A / Ra
   cGasolina := G / Rg

   if cAlcool < cGasolina {
    fmt.Println("A")
   }else {
    fmt.Println("G")
   }
}
