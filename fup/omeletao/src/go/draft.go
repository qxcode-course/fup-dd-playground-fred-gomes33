package main
import "fmt"
func main() {
    var num1, num2, num3, num4 int
    fmt.Scan(&num1, &num2, &num3, &num4)

    maior := num1
    if num2 > maior {
        maior = num2
    }
    if num3 > maior {
        maior = num3
    }
    if num4 > maior {
        maior = num4
    }
    fmt.Println(maior)
}
