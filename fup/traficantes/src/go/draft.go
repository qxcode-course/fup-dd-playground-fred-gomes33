package main
import "fmt"
func main() {
	var frase string
	var caractere byte

	for {
		n, _ := fmt.Scanf("%c", &caractere)
		if n != 1 || caractere == '\n' {
			break
		}
		if caractere != '\r' {
			frase += string(caractere)
		}
	}

	var antiga, nova string
    fmt.Scan(&antiga, &nova)
	tamFrase := len(frase)
	tamAntiga := len(antiga)

	for i := 0; i < tamFrase; {
		if i+tamAntiga <= tamFrase && frase[i:i+tamAntiga] == antiga {
			fmt.Print(nova)
			i += tamAntiga
		} else {
			fmt.Printf("%c", frase[i])
			i++
		}
	}
	fmt.Println()
}