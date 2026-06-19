package main
import "fmt"
func main() {
    nomes := make(map[string]int)
    
    for{
    palavras := ""
    fmt.Scan(&palavras)

    if palavras == "pare" { 
        break
    }
        nomes[palavras] += 1
}
    fmt.Println(nomes)
}