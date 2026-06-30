package main
import "fmt"
func ehVogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func main() {
	var palavra, resultado string

	if n, _ := fmt.Scan(&palavra); n == 1 {
		resultado = palavra
	}

	for n, _ := fmt.Scan(&palavra); n == 1; {
		tamRes := len(resultado)

		if tamRes > 0 && ehVogal(resultado[tamRes-1]) && ehVogal(palavra[0]) {
			for len(resultado) > 0 && ehVogal(resultado[len(resultado)-1]) {
				resultado = resultado[:len(resultado)-1]
			}
			for len(palavra) > 0 && ehVogal(palavra[0]) {
				palavra = palavra[1:]
			}
			resultado += palavra
		} else {
			resultado += " " + palavra
		}
	}

	fmt.Println(resultado)
}