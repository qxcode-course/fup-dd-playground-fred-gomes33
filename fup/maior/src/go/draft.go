package main
import "fmt"
func main() {
   var c1 float64
   var esc string
   var valr float64 
   fmt.Scan(&c1, &esc, &valr)

    if c1 == valr{
        fmt.Println("primeiro")
    } else if esc == "m"{
        if valr < c1{
            fmt.Println("segundo")
        } else {
            fmt.Println("primeiro")
        }
    } else if esc == "M"{
        if valr > c1{
            fmt.Println("segundo")
        } else {
            fmt.Println("primeiro")
        }
    }
}
