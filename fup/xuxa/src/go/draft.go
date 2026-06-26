package main
import "fmt"
func main() {
    var texto string
    for {
        var caractere string
        fmt.Scanf("%s", &caractere)
        if caractere == "\n"{
            break
        }
        texto += caractere
    }

    runes := []rune(texto)
    n := len(runes)

    for i:= 0; i < n/2; i++{
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    fmt.Println(string(runes))
}