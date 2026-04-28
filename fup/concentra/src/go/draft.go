package main
import "fmt"
func main() {
   var j1, j2 int
   fmt.Scan(&j1, &j2)

   fmt.Print("[ ")
   for i := 0; i <= 10; i++ {
      if i == j1 || i == j2 {
         continue
      }
      if i == 10 {
         fmt.Print(i)
      } else {
         fmt.Print(i, " ")
      }
   }
   fmt.Print(" ]")
}
