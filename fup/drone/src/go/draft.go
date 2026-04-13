aapackage main
import "fmt"
import "sort"
func main() {

var A, B, C, H, L int
fmt.Scan(&A, &B, &C, &H, &L)

janela := []int{H, L}
sort.Ints(janela)

caixa := []int{A, B, C}
for i := 0; i < 3; i++ {
    for j := i + 1; j < 3; j++ {
        face := []int{caixa[i], caixa[j]}
        sort.Ints(face)
        if face[0] <= janela[0] && face[1] <= janela[1] {
            fmt.Println("S")
            return
    }   
    }
}
fmt.Println("N")
}  
