package main
import "fmt"
func main() {
   var n int
   fmt.Scan(&n)

   soldados := 0
   rebeldes := 0

   for i := 0; i < n; i++{
    var x int
    fmt.Scan(&x)
    if x%2 == 0{
        rebeldes += x
    }else {
        soldados += x
    }
   }
   if soldados > rebeldes{
    fmt.Println("soldados")
   }else if rebeldes > soldados{
    fmt.Println("rebeldes")
   }else{
    fmt.Println("empate")
   }
}