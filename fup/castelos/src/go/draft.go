package main

import "fmt"

func main() {
	var n int
    fmt.Scan(&n)

    ehQuadrado := false
    for i := 1; i*i <= n; i++ {
        if i*i == n {
            ehQuadrado = true
            break
        }
    }

    if ehQuadrado {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
