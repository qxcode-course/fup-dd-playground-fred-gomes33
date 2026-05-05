package main
import "fmt"
func main() {
    var C int
    fmt.Scan(&C)
    
    pass := 0
    for{
        var M int
        _, err := fmt.Scan(&M)
        if err != nil{
            break
        }
        pass += M

       if pass >= 2*C {
            fmt.Println("hora de partir")
            break
        } else if pass == 0 {
            fmt.Println("vazio")
        } else if pass < C {
            fmt.Println("ainda cabe")
        } else if pass == C {
            fmt.Println("lotado")
        } else if pass > C {
            fmt.Println("lotado")
        }
    }
}
