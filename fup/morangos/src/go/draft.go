package main
import "fmt"
func main() {
  //ter a melhor área
    var c, l, c2, l2 int
    fmt.Scan(&c, &l, &c2, &l2)

    area1 := c * l
    area2 := c2 * l2

    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
}
