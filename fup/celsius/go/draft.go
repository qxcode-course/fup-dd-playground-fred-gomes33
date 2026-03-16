package main
import "fmt"
func main() {
    var celsius, fahrenheit float64
   fmt.Scan(&celsius, &fahrenheit)
   fahrenheit = (celsius * 9/5) + 32
    fmt.Println(fahrenheit, "%.2f\n",)
}
