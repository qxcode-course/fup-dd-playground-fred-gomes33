package main
import "fmt"
func main() {
    var im, if1, if2, if3, maior int
    fmt.Scan(&im, &if1, &if2, &if3, &maior)

    if3 = im - (if1 + if2)
    maior = if1

    if if2 > maior {
        maior = if2
    }

    if if3 > maior {
        maior = if3
    }

    fmt.Println(maior)
}
