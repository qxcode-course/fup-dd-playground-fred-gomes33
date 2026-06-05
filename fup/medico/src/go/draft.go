package main
import "fmt"
func main() {
    var qnt int
    fmt.Scan(&qnt)

    seq := make([]int, qnt)

    for i := 0; i < qnt; i++ {
        fmt.Scan(&seq[i])
    }

    cont := 0
    
    for i := 0; i < len(seq); i++ {
        if seq[i] == 0 {
            tup := false
            if i > 0 && seq[i-1] == 1 {
                tup = true
            }
            if i < len(seq)-1 && seq[i+1] == 1 {
                tup = true
            }
            if tup == false {
                cont++
            }
        }
        
    }
fmt.Println(cont)
}