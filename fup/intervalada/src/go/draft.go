package main
import "fmt"
func main() {
    var n, min, max int
    if _, err := fmt.Scan(&n, &min, &max); err != nil{
        return
    }

    count := 0

    for i := 0; i < n; i++{
        var num int
        fmt.Scan(&num)
        
        if num >= min && num <= max {
            count++
        }
    }
    fmt.Println(count)
}