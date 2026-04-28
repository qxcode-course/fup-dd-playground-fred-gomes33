package main
import "fmt"
func main() {
    var pedra int
    var inicio string
    fmt.Scan(&pedra, &inicio)

    pe := inicio

    fmt.Print("[ ")
    for i := 0; i <= 10; i++ {
        if i == pedra{
            continue 
        }
        if i == 10{
            fmt.Print("ceu ")
            break
        }    
        fmt.Printf("%d%s ", i, pe)
        if pe == "d"{
            pe = "e"
        } else{
            pe = "d"
        }
  }
  fmt.Print("]\n")
}
