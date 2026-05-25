package main
import "fmt"
func main() {
   var qtd int
    fmt.Scan(&qtd)

    c := 0
    l := 0
    m := 0
    t := 0

    for i := 0; i < qtd; i++ {
        var sabor, turno string
        fmt.Scan(&sabor, &turno)

        switch sabor {
        case "c":
            c++
        case "l":
            l++
        }

        switch turno {
        case "m":
            m++
        case "t":
            t++
        }
    }

    if c > l{
        fmt.Println("c")
    }else if l > c{
        fmt.Println("l")
    }else{
        fmt.Println("empate")
    }
    if m > t {
        fmt.Println("m")
    } else if t > m {
        fmt.Println("t")
    } else {
        fmt.Println("empate")
    }
}
