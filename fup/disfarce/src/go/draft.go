package main
import "fmt"

func toLower(c byte) byte{
    if c >= 'A' && c <= 'Z'{
        return c + 32
    }
    return c
}
func main() {
    var x int
    fmt.Scan(&x)

    for i := 0; i < x; i++{
        var ultron, pessoa string
        fmt.Scan(&ultron, &pessoa)

        correspondencias := 0

        for j := 0; j < len(pessoa); j++{
            pc := toLower(pessoa[j])
            for k := 0; k < len(ultron); k++{
                uc  := toLower(ultron[k])
                if pc == uc{
                    correspondencias++
                    break
            }
        }
    }

    porcentagem := float64(correspondencias) / float64(len(pessoa)) * 100

    if porcentagem == 100{
        fmt.Println("chefe")
    }else if porcentagem > 50{
        fmt.Println("ultron")
    }else{
    fmt.Println("pessoa")
    }
 }
}
