package main
import "fmt"
func main() {
   var n int
   if _, err := fmt.Scan(&n); err != nil{
    return
   }

   array := make([]int, n)
   for i := 0; i < n; i++{
    fmt.Scan(&array[i])
   }
   
   casais := 0

   for i := 0; i < n; i++{
    if array[i] == 0{
        continue
    }
     for j := i + 1; j < n; j++{
       if array[i] == -array[j]{
         casais++
       
        array[i] = 0
        array[j] = 0
        break
        }
    }
}
    fmt.Println(casais)
}