package main
import "fmt"

func FatorarPrimos(n int){
    fator := 2
    contagem := 0

    for n != 1{
        if n%fator == 0{
            n /= fator
            contagem++
        }else{
            if contagem > 0{
                fmt.Printf("%d %d\n", fator, contagem)
            }
            fator++
            contagem = 0
        }
    }
    if contagem > 0{
        fmt.Printf("%d %d\n", fator, contagem)
    }
}

func main(){
    var n int
    fmt.Scan(&n)
    FatorarPrimos(n)
}
