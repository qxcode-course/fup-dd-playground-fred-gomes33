package main
import "fmt"
func main() {
   corretas := []rune{'d', 'a', 'c', 'd'}
   respostas := make([]rune, 4)

   for i := 0; i < 4; i++{
    var r string
    fmt.Scan(&r)
    respostas[i] = rune(r[0])
   }
 
   acertos := 0
   for i := 0; i < 4; i++{
    if respostas[i] == corretas[i]{
        acertos++
    }
   }
   
    switch acertos {
    case 0:
        fmt.Println("Nunca assistiu")
    case 1:
        fmt.Println("Ja ouviu falar")
    case 2:
        fmt.Println("Interessado no assunto")
    case 3:
        fmt.Println("Fa")
    case 4:
        fmt.Println("Super Fa")
    }
}
