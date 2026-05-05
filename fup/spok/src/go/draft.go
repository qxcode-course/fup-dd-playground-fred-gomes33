package main
import "fmt"

func Pali(n int) int {
    ori := n
    inv := 0

    for n > 0 {
        digi := n % 10
        inv = inv*10 + digi
        n = n / 10
    }
    
    if inv == ori{
        return 1
    }
    return 0
}

func main() {
    var id int
    fmt.Scan(&id)
    fmt.Println(Pali(id))
}
