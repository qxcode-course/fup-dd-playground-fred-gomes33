package main
import "fmt"
func main() {
    var idf1, qtnf int
    fmt.Scan(&idf1, &qtnf)

    for i := 0; i < qtnf; i++ {
        fmt.Println(idf1 + i*2)
    
    }
}
