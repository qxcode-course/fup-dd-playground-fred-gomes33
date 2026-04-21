package main
import "fmt"
func main() {
   var nl, nc int
   fmt.Scan(&nl, &nc)

   if (nl+nc)%2 == 0{
    fmt.Println(1)
   } else {
    fmt.Println(0)
   }
}
