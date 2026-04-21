package main
import "fmt"
func main() {
   var n1, n2 int
   fmt.Scan(&n1, &n2)

    if (n1 % 3 == 0 && n2 % 3 == 0) || (n1 % 5 == 0 && n2 % 5 == 0) {
            fmt.Println("sim")
        } else {
            fmt.Println("nao")
        }
}
