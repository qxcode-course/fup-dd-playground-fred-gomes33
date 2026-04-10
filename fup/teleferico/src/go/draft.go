package main
import "fmt"
import "math"

func main() {
    var C, A int
    fmt.Scan(&C, &A)

    capacidadeAlunos := C - 1
    viagens := int(math.Ceil(float64(A) / float64(capacidadeAlunos)))

    fmt.Println(viagens)
}
