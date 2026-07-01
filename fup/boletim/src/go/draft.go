package main
import "fmt"
func main() {
    soma := 0
    var nota int

    for i := 0; i < 2; i++{
        for j := 0; j < 3; j++{
        fmt.Scan(&nota)
        soma += nota
    }
}
    fmt.Println(soma)
}