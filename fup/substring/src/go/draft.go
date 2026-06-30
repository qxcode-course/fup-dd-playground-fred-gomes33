package main
import "fmt"
func main() {
    var texto string
    var caractere byte

    for {
        var n int
        var err error
        n, err = fmt.Scanf("%c", &caractere)
        if n != 1 || err != nil || caractere == '\n' {
            break
        }
        if caractere != '\r' {
            texto += string(caractere)
        }
    }
    var inicio, qtd int
    fmt.Scan(&inicio, &qtd)
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