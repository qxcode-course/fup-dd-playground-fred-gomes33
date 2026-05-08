package main
import "fmt"
func main() {
 var h, m, d int
 var s string

   fmt.Scan(&h, &m, &d, &s)

   
    minutosIniciais := h*60 + m

    minutosCaminhados := d * 10

    var minutosFinais int
    if s == "H" {
        minutosFinais = minutosIniciais + minutosCaminhados
    } else {
        minutosFinais = minutosIniciais - minutosCaminhados
    }

    minutosFinais = ((minutosFinais % 720) + 720) % 720

    hora := minutosFinais / 60
    minuto := minutosFinais % 60

    fmt.Printf("%02d %02d\n", hora, minuto)
}


