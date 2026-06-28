package main
import "fmt"
import "unicode"
func main() {
   var caractere rune
   if _, err := fmt.Scanf("%c", &caractere); err != nil{
    return
   }

   if unicode.IsLower(caractere){ //indentifica se a letra é minu.
    caractere = unicode.ToUpper(caractere)
   } else if unicode.IsUpper(caractere) { //indentifica se a letra é maiu.
    caractere = unicode.ToLower(caractere)
   }
    fmt.Printf("%c\n", caractere)
}