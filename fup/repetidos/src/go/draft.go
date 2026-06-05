package main
import "fmt"
func main() {
     
    var P, N int
    fmt.Scan(&P, &N)

    count := 0
    for i := 0; i < N; i++ {
        var x int
        fmt.Scan(&x)
        if x == P {
            count++
        }
    }

    fmt.Println(count)
}