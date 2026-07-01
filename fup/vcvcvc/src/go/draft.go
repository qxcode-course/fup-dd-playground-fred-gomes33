package main
import "fmt"
func main() {
    var caractere byte

    for {
        n, err := fmt.Scanf("%c", &caractere)
        if err != nil || n != 1 || caractere == '\n' {
            break
        }

        if caractere == '\r' {
            continue
        }

        if caractere == ' '{
            fmt.Print(" ")
        } else {
            letra := caractere
            if letra >= 'A' && letra <= 'Z' {
                letra = letra + 32
            }

            if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' {
                fmt.Print("v")
            } else if letra >= 'a' && letra <= 'z' {
                fmt.Print("c")
            }
        }
    }
    fmt.Println()
}

