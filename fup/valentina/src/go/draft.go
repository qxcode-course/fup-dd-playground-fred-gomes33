package main
import "fmt"
func main() {
    var c1, c2, op string
    n, err := fmt.Scan(&c1, &op, &c2)
    if err != nil || n != 3 {
        return
    }

    v1 := int(c1[0] - 'a')
    v2 := int(c2[0] - 'a')
    var resultado int

    if op == "+"{
        resultado = (v1 + v2) % 26
    }else {
        resultado = (v1 - v2 + 26) % 26
    }
    fmt.Printf("%c\n", 'a'+resultado)
}