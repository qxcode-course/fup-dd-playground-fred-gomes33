package main
import "fmt"
func main() {
    var atual, anterior string
    ordenado := true

    if n, _ := fmt.Scan(&anterior); n == 1 {
        for n, _ := fmt.Scan(&atual); n == 1; n, _ = fmt.Scan(&atual) {
            if atual < anterior {
                ordenado = false
            }
            anterior = atual
        }
    }

    if ordenado {
        fmt.Println("sim")
    }else{
        fmt.Println("nao")
    }
}