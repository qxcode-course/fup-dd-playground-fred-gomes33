package main
import "fmt"
func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil{
        return
    }

    ledsPdigito := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

    for i := 0; i < n; i++{
        var v string
        fmt.Scan(&v)

        totaldleds := 0
        for _, digito := range v{
            totaldleds += ledsPdigito[digito-'0']
        }
        fmt.Printf("%d leds\n", totaldleds)
    }
}