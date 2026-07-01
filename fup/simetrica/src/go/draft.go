package main
import "fmt"
func main() {
    var mat [3][3]int
    for i := 0; i < 3; i++{
        for j := 0; j < 3; j++{
            fmt.Scan(&mat[i][j])
        }
    }

   simetrica := true

   for i := 0; i < 3; i++{
    for j := 0; j < 3; j++{
        if mat[i][j] != mat[j][i]{
            simetrica = false
            break
        }
    }
    if !simetrica{
        break
    }
}
    if simetrica {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}