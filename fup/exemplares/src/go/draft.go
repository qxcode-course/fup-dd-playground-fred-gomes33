package main
import "fmt"
import "sort"
func contains(array []int, valor int) bool {
    for _, elem := range array {
        if elem == valor {
            return true
        }
    }
    return false
}

func separador(array []int) []int {
    especies := make([]int, 0)
    for _, animal := range array{
        if !contains(especies, animal){
            especies = append(especies, animal)
        }
    }
    return especies
}
func main(){
    var qtd int
    fmt.Scan(&qtd)
    animais := make([]int, qtd)
    for i :=range animais {
        fmt.Scan(&animais[i])
    }
    especies := separador(animais)
    sort.Ints(especies)

    for i, elem := range especies {
        if i == len(especies) - 1 {
            fmt.Println(elem)
            break
        }else{
            fmt.Print(elem, " ")
        }
    }
}