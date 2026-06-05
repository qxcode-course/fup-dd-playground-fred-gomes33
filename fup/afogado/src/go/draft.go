package main
import "fmt"
func main() {
    var P, E int
    fmt.Scan(&P, &E)

    for forcaini := 1; forcaini <= P; forcaini++ {
        for forcaini2 := 1; forcaini2 <= E; forcaini2++ {
            if (forcaini*forcaini2) == P {
                if (forcaini + forcaini2) == E {
                    fmt.Printf("%d %d\n", forcaini, forcaini2)
                }
            }
        }
    }
}