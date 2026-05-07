package main
import "fmt"
func main() {
   var angulo int
   fmt.Scan(&angulo)

   angulo = ((angulo % 360) + 360) % 360
   fmt.Println(angulo)
}
