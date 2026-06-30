package main
import "fmt"
func main() {
    var palavra string
    var vogais, consoantes []rune

    for {
        if _, err := fmt.Scan(&palavra); err != nil {
            break
        }
        for _, letra := range palavra {
            switch letra {
            case 'a', 'e', 'i', 'o', 'u':
                vogais = append(vogais, letra)
            default:
                consoantes = append(consoantes, letra)
            }
        }
    }
    fmt.Println(string(vogais))
    fmt.Println(string(consoantes))
}
