package main
import "fmt"
func main() {
   /*
   L -> atacante que lança a bola
   R -> jogador atacante que recebe a bola
   D -> jogador defensor
   regra -> seu campo de ataque  || linha divisória do meio do campo
   Campo -> 100 m
   */
   var L, R, D int
   fmt.Scan(&L,&R, &D)
   
    if R > 50 && L < R && R > D{
        fmt.Println("S")
    }else{
        fmt.Println("N")
     }
}
