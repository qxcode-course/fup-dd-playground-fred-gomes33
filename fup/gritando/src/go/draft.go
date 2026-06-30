package main
import "fmt"
func main() {
    var palavra string
    primeira := true

    for {
        n, err := fmt.Scan(&palavra)
        if n != 1 || err != nil {
            break
        }

        if !primeira {
            fmt.Print(" ")
        }

        primeira = false

        for _, letra := range palavra{
            if letra >= 'a' && letra <= 'z'{
                fmt.Printf("%c", letra-32)
            }else if letra >= 'A' && letra <= 'Z'{
                fmt.Printf("%c", letra+32)
            }else{
                fmt.Printf("%c", letra)
            }
        }
    }
    fmt.Println()
}