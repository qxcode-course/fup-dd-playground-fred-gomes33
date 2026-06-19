package main
import "fmt"
func main() {
    var total, qtd int
   fmt.Scan(&total, &qtd)

   figurinhas := make([]int, qtd)
    for i:= 0; i < qtd; i++{
        fmt.Scan(&figurinhas[i])
    }
    repetidas := []int{}
    for i := 1; i < qtd; i++{
        if figurinhas[i] == figurinhas[i-1]{
            repetidas = append(repetidas, figurinhas[i])
        }
    }

    faltando := []int{}
    idx := 0
    for num := 1; num <= total; num++{
       for idx < qtd && figurinhas[idx] < num{
        idx++
       }
       if idx >= qtd || figurinhas[idx] != num{
        faltando = append(faltando, num)
       }
    }
    fmt.Print("[ ")
    for i, r := range repetidas {
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(r)
    }
    fmt.Println(" ]")

    fmt.Print("[ ")
    for i, f := range faltando{
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(f)
    }
    fmt.Println(" ]")
}
