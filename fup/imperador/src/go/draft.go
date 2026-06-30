package main
import "fmt"
func main() {
    var n int
    _, err := fmt.Scan(&n)
    if err != nil{
        return
    }
    arena := make([][]string, n)
    leaoL, leaoC := -1, -1

    for i := 0; i < n; i++{
        arena[i] = make([]string, n)
        for j := 0; j < n; j++{
            fmt.Scan(&arena[i][j])
            if arena[i][j] == "L"{
                leaoL = i
                leaoC = j
            }
        }
    }
    pontosG := 0
    pontosC := 0

    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            if i == leaoL || j == leaoC{
                continue
            }

            switch arena[i][j] {
            case "G":
                pontosG += 2
            case "C":
                if  i+j == n-1 {
                    pontosC += 2
                }else {
                    pontosC += 1
                }
            }
        }
}
if pontosG > pontosC{
        fmt.Println("Gladiadores")
}else if pontosC > pontosG {
        fmt.Println("Condenados a morte")
    }else {
        fmt.Println("Ninguem")
    }
}