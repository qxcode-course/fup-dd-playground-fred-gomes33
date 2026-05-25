package main

import (
	"fmt"
	"math"
)

func main() {
    var chutechico, chutecebolinha, qtdanimais int
    fmt.Scan(&chutechico, &chutecebolinha, &qtdanimais)

    totalPatas := 0
    for i := 0; i < qtdanimais; i++ {
        var animal string
        fmt.Scan(&animal)
        switch animal {
        case "v":
            totalPatas += 4
        case "g":
            totalPatas += 2
        case "c":
            totalPatas += 4
        }
    }
    fmt.Print(totalPatas)

    distChico := int(math.Abs(float64(chutechico - totalPatas)))
    distCebolinha := int(math.Abs(float64(chutecebolinha - totalPatas)))

    if distChico < distCebolinha {
        fmt.Println("\nChico Bento")
    } else if distCebolinha < distChico {
        fmt.Println("\nCebolinha")
    } else {
        fmt.Println("\nempate")
    }
}