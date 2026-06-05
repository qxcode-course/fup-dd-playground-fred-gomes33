package main
import "fmt"
import "math"
func main() {
   var num int
   fmt.Scan(&num)

   primo:= 1

   for i := 2; i <= int(math.Sqrt(float64(num))); i++{
    if num % i == 0{
        primo = 0
        break
    }
   }
    fmt.Println(primo)
}