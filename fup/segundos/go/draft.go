package main
import "fmt"

func main() {
    var temp, hor, min, seg, restt, rests int64
    fmt.Scan(&temp)
    hor = temp / 3600 
    restt = temp % 3600
    min = restt / 60
    rests = restt % 60
    seg = rests % 60
    fmt.Printf("%d:%d:%d\n", hor, min, seg)
}
