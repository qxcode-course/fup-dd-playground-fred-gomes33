package main

import "fmt"
func main() {
   var qtd, cont, res int
   fmt.Scan(&qtd)

   blocos := make([]int, qtd)
   for i := range blocos {
       fmt.Scan(&blocos[i])
   }
   res = blocos[0]
   for _, elem := range blocos{
    if elem > res{
        if elem - 1 > res{
            cont += 1
            res = elem
            continue
        }else{
            res = elem
            continue
        }
    }else if elem < res {
        if elem + 1 < res{
            cont += 1
            res = elem
            continue
        }else {
            res = elem
            continue
        }
    }
}
   
    fmt.Println(cont)
}