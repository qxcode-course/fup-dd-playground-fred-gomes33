package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil{
        return
    }    
    var alunos []int
    var servidores []int

    for i := 0; i < n; i++{
        var numero int
        fmt.Scan(&numero)

        if numero%2 == 0{
            servidores = append(servidores, numero)
            }else{
                alunos = append(alunos, numero)
            }
        }
        imprimirresultado(alunos)
        imprimirresultado(servidores)
    }
func imprimirresultado(array []int){
    fmt.Print("[")
    for _, num := range array{
        fmt.Printf(" %d", num)
    }
    fmt.Println(" ]")
}