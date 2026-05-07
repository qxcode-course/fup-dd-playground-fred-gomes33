package main
import "fmt"
func main() {
   var tipo string
   var F int
   fmt.Scan(&tipo, &F)

   var T int
   if tipo == "b"{
    T = 20
   }else if tipo == "c"{
    T = 18
   }
   
   P := ((F * T ) - 80) / 10

   if P < 150{
    fmt.Println("Fraco, nem passou")
   }else if P < 180{
    fmt.Println("Perfeito")
   }else if P < 210{
    fmt.Println("Satisfeito")
   }else{ 
    fmt.Println("Muito forte, bola fora")
 }
}