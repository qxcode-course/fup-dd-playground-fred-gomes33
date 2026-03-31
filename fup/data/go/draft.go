package main
import "fmt"
func main() {
   var hor, min, dia, mes, ano int
   fmt.Scan(&hor, &min, &dia, &mes, &ano)
    ano = ano % 100

    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", hor, min, dia, mes, ano)
    
}
