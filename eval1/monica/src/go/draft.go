package main
import "fmt"
func main() {
    var M, f1, f2 int
    fmt.Scan(&M, &f1, &f2)

    f3 := M - (f1 + f2)

    maior := f1
    if f2 > maior {
        maior = f2
    }
    if f3 > maior {
        maior = f3
    }
        fmt.Println(maior)
}
