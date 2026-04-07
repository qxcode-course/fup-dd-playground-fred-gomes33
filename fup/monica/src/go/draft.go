package main
import "fmt"
func main() {
    var im, if1, if2, if3 int 
    fmt.Scan(&if1, &if2, &if3, &im)

    if3 = if1 + if2 - im
    if3 = if3 * 2
    fmt.Println(if3)
}
