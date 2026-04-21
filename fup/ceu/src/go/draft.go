package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Print("[ ")
    for i := 0; i <= 10; i++{
        //n onde a pedra caiu, aí você pula
        if i == n{
            continue
        }
        if i == 10 {
            fmt.Print("ceu ")
        }else{
            fmt.Printf("%d ", i)
        }
    }    
    
    fmt.Println("]")
}
