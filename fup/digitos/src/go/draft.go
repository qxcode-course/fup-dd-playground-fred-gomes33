package main
import "fmt"
import "strconv"
func main() {
    var A, B int
    fmt.Scan(&A, &B)

    contato := strconv.Itoa(B)
    digito := strconv.Itoa(A)

    count := 0
    for _, ch := range fmt.Sprintf(contato) {
        if string(ch) == digito {
            count++
        }
    }
    fmt.Println(count)
}