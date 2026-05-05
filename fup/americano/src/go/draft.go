package main
import "fmt"
func main() {
    var j1, j2, j3, j4 int
    fmt.Scan(&j1, &j2, &j3, &j4)

    soma := j1 + j2 + j3 + j4

    if soma == 0{
        fmt.Println("nenhum")
        return
    }
    
    switch soma % 4{
    case 1: 
        fmt.Println("jog1")
    case 2:
        fmt.Println("jog2")
    case 3:
        fmt.Println("jog3")
    case 0:
        fmt.Println("jog4")
    }
}
