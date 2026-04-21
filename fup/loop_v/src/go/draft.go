package main
import "fmt"
func main() {
   var A, B int
   fmt.Scan(&A, &B)

   fmt.Printf("[ ")
   i := A

   for {
    if i == B{
        break
    }
    if i%2 == 0{
        i++
        continue
    }
    fmt.Printf("%d ", i)
    i++
   }
   fmt.Printf("]\n")
}