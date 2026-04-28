package main
import "fmt"
func main() {
   var ni, nf int
   fmt.Scan(&ni, &nf)

   for i := ni; i <= nf; i++ {
        if i %3 == 0 && i %5 == 0{
            fmt.Println("zigzag")
        } else if i %3 == 0{
            fmt.Println("zig")
        }else if i %5 == 0{
            fmt.Println("zag")
        }else{
            fmt.Println(i)
        }
   }
}
