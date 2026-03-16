package main
import "fmt"
func main() {
    var num1, num2 float32
    fmt.Scan(&num1, &num2)
    var s float32 = ((num1 + num2)/2)
    fmt.Printf("%.1f\n", s)

}
