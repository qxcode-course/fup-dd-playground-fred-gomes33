package main
import "fmt"
func main() {
   var mat [3][3]int

   for i := 0; i < 3; i++{
        for j := 0; j < 3; j++{
            fmt.Scan(&mat[i][j])
        }
    }

base := mat[0][0] + mat[0][1] + mat[0][2]
magico := true

for i := 0; i < 3; i++{
    if mat[i][0] + mat[i][1] + mat[i][2] != base{
        magico = false
    }
}

for j := 0; j < 3; j++{
    if mat[0][j] + mat[1][j] + mat[2][j] != base{
        magico = false
    }
}

    if mat[0][0] + mat[1][1] + mat[2][2] != base{
        magico = false
}
    if mat[0][2]+mat[1][1]+mat[2][0] != base {
		magico = false
}
    if magico {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}