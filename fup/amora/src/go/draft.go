package main
import "fmt"
func main() {
    var frase string
    var caractere byte

    for {
        n, err := fmt.Scanf("%c", &caractere)
        if n != 1 || err != nil || caractere == '\n' {
            break
        }
        if caractere != '\r' {
            frase += string(caractere)
        }
    }
    var trecho string
    fmt.Scan(&trecho)

    count := 0
    tamFrase, tamTrecho := len(frase), len(trecho)


    for i := 0; i <= tamFrase-tamTrecho; i++{
        if frase[i:i+tamTrecho] == trecho{
            count++
        }
    }
    fmt.Println(count)
}