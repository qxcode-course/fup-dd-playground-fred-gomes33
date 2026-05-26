package main
import "fmt"

// 0 -> ouros; 1 -> copas; 2 -> paus; 3 -> espadas
type Carta struct {
    numero int
    naipe int
}

func toString(carta Carta) string { // Carta{1,1} -> A ♥
    var symbols = []string{"♦", "♥", "♣", "♠"}
    naipe := symbols[carta.naipe]
    var numeros = []string{"", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
    return fmt.Sprintf("%s%s", numeros[carta.numero], naipe)
}

func main() {
    for numero := 1; numero <= 13; numero++ {
        for naipe := 0; naipe <= 3; naipe++ {
            fmt.Printf("%v ", toString(Carta{numero, naipe}))
        }
    }
    fmt.Println("")
}

func criarBaralho() []Carta {
    baralho := make([]Carta, 0, 52)
    for naipe := 0; naipe <= 3; naipe++ {
        for numero := 1; numero <= 13; numero++ {
            baralho = append(baralho, Carta{numero, naipe})
        }
    }
    return baralho
}