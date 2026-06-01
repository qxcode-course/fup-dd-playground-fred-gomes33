package main

import "fmt"

// 0 "ouros", 1 "copas", 2 "paus", 3 "espadas"
type Carta struct {
	numero int
	naipe  int
}

type Deck struct {
	cartas []Carta
}

func (deck *Deck) adicionarCarta(carta Carta) {
	deck.cartas = append(deck.cartas, carta)
}

// método
func (carta Carta) toString() string { // Carta{12, 2} -> Q♣
	var symbols = []string{"♦", "♥", "♣", "♠"}
	var numeros = []string{"", " A", " 2", " 3", " 4", " 5", " 6", " 7",
		" 8", " 9", "10", " J", " Q", " K"}
	return fmt.Sprintf("%v%v", numeros[carta.numero], symbols[carta.naipe])
}

func criarBaralho() Deck {
	baralho := make([]Carta, 0, 52)
	for naipe := 0; naipe <= 3; naipe++ {
		for num := 1; num <= 13; num++ { //1
			baralho = append(baralho, Carta{numero: num, naipe: naipe})
		}
	}

	return Deck{baralho}
}

func (deck *Deck) puxarCarta() Carta {
	if len(deck.cartas) == 0 {
		panic(1)
	}
	carta := deck.cartas[len(deck.cartas)-1]
	deck.cartas = deck.cartas[:len(deck.cartas)-1]
	return carta
}

func main() {
	baralho := criarBaralho()
	jogador := Deck{}
	jogador.adicionarCarta(baralho.puxarCarta())
	jogador.adicionarCarta(baralho.puxarCarta())
    for _, carta := range jogador.cartas {
        fmt.Println(carta.toString())
    }

}
