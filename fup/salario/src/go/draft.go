package main
import "fmt"
func main() {
   var salt, nov float64
   fmt.Scan(&salt, &nov)
   
    if salt <= 1000{
            nov = salt + (salt * 0.20)
        }else if salt > 1000 && salt <= 1500{
            nov = salt + (salt * 0.15)
        }else if salt > 1500 && salt <= 2000{
            nov = salt + (salt * 0.10)
        }else if salt > 2000{
            nov = salt + (salt * 0.05)
        }
    fmt.Printf("%.2f\n", nov)
}
