package main
import "fmt"
func main() {
    var soma int
    fmt.Scan(&soma)

    if soma == 0{
        fmt.Println("joguem de novo")
        return
    }

    i := (soma - 1) % 26
    letra := rune('a') + rune(i)
    
    fmt.Printf("%c\n", letra)
}
