package main

import "fmt"

type Pessoa struct {
	Nome        string
	Idade       int
	Raca        string
	Sexualidade string
	numfav      int
}

func main() {

	/* Array := []string{"Maçã", "Banana", "Abacate"}
	   Array = append(Array, "Uva") //-> adiciona Uva no final


	   fmt.Println(Array)*/

	/*Array := []string{"Cenoura", "Alface", "Beterraba"}
	  Array = append(Array, "Melão")
	  for _, Compras := range Array{
	      fmt.Println("Comprei:", Compras)
	  }*/
	// 1ª Pessoa

	pessoa1 := Pessoa{"Felipe", 18, "Parda", "Gay", 7}
	fmt.Println(pessoa1.Nome,"\n",pessoa1.numfav,pessoa1.Raca)

}
