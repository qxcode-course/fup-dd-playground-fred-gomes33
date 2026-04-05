package main
import "fmt"
func main() {
    var num1, num2 int
    var op string
    var resultado int
    fmt.Scan(&num1, &num2, &op)

    switch op {
case "+":
        resultado = num1 + num2
    case "-":
        resultado = num1 - num2
    case "*":
        resultado = num1 * num2
    case "/":
        if num2 != 0 {
            resultado = num1 / num2 // divisão inteira
        } else {
            resultado = 0
        }
    }
    fmt.Println(resultado)
}
