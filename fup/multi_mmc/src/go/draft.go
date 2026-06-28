package main
import "fmt"

func mdc(a, b int) int{
    for b != 0{
        a, b = b, a%b
    }
    return a
}

func mmcArray(numeros []int) int{
    if len(numeros) == 0{
        return 0
    }
    resultado := numeros[0]
    for _, num := range numeros[1:]{
        if resultado == 0 || num == 1{
            return 0
        }

        resultado = (resultado / mdc(resultado, num)) * num
    }
    return resultado
}
func main() {
   var n int
   if _, err := fmt.Scan(&n); err != nil || n <= 0 {
        return
   }

   Array := make([]int, n)
   for i := 0; i < n; i++{
    fmt.Scan(&Array[i])
   }
    fmt.Println(mmcArray(Array))
}