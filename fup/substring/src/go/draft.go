package main
import "fmt"
func main() {
    var texto string
    var inicio, qtd int

    if _, err := fmt.Scan(&texto, &inicio, &qtd); err != nil{
        return
    }

    if inicio < 0 || inicio >= len(texto) || qtd <= 0{
        fmt.Println("")
        return
    }

    fim := inicio + qtd
    if fim > len(texto){
        fim = len(texto)
    }
    fmt.Println(texto[inicio:fim])
}