package main
import "fmt"
func main() {
    var n, x int
    if _, err := fmt.Scan(&n, &x); err != nil{
        return
    }

    operarios := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&operarios[i])
    }

    for j := 0; j < x; j++{
        var grito int
        fmt.Scan(&grito)

        indiceChefe := 1

        for i, v := range operarios{
            if v== grito || v == -grito{
                indiceChefe = i
                break
            }
        }

        if indiceChefe == -1{
            continue
        }

        if indiceChefe-1 >= 0{
            operarios[indiceChefe-1] = -operarios[indiceChefe-1]
        }

        if indiceChefe+1 < n{
           operarios[indiceChefe+1] = -operarios[indiceChefe+1]
        }
    }
    fmt.Printf("%v\n", operarios)
}