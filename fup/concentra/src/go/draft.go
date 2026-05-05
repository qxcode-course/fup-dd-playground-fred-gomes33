package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    fmt.Print("[ ")
    for i := A; i <= B; i++ {
        j := A + B - i
        fmt.Print(i, " ", j, " ")
        if i == j {
            break
        }
    }

   
    for i := B; i >= A; i-- {
        j := A + B - i
        fmt.Print(i, " ", j, " ")
        if i == j {
            break
        }
    }

    fmt.Println("]")
}