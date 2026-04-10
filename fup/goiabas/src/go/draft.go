package main
import "fmt"
import "math"

func main() {
   var C, ban, goi, man, tot, via int
    fmt.Scan(&C, &ban, &goi, &man, &tot, &via)  

    tot = ban + goi + man
    via = int(math.Ceil(float64(tot) / float64(C)))

    fmt.Println(via)
}
