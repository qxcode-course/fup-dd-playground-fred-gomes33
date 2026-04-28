# Aula 2 | Vetores

package main

import (
	"fmt"
	"slices"
	"strconv"
)

func filtrar_impares(nums []int) []int {
	impares := make([]int, 0, len(nums))
	for _, elem := range nums {
		if elem%2 == 1 {
			impares = append(impares, elem)
		}
	}
	return impares
}

func index(nums []int, valor int) int {
	for i, elem := range nums {
		if elem == valor {
			return i
		}
	}
	return -1
}

func contains(nums []int, valor int) bool {
	for _, elem := range nums {
		if elem == valor {
			return true
		}
	}
	return false
}

func count(nums []int, valor int) int {
	contador := 0
	for _, elem := range nums {
		if elem == valor {
			contador += 1
		}
	}
	return contador
}

func separar_figurinhas(montante []int) ([]int, []int) { //tupla
	album := make([]int, 0, len(montante))
	repet := make([]int, 0, len(montante))
	for _, fig := range montante {
		if !contains(album, fig) {
			album = append(album, fig)
		} else {
			repet = append(repet, fig)
		}
	}
	return album, repet
}

func main() {
	var montante []int = make([]int, 0, 1)
	fmt.Println(montante, len(montante), cap(montante))
	montante = append(montante, 7, 3, 2, 1, 9, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 5, 7)
	// album: 1, 2, 3, 4, 5, 7
	// trocar: 4, 3, 2, 1, 2, 5
	num, err := strconv.Atoi("32432")
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err)
	}
	album, repet := separar_figurinhas(montante)

	slices.Sort(repet)
	fmt.Println(album)
	fmt.Println(repet)


}

func main() {

## Como usar os rascunhos

- A chave do seu rascunho é o nome da pasta.
- O título do seu rascunho é carregado a partir da primeira linha desse arquivo \#
- Tudo que você fizer nos rascunhos também será rastreado pelo tko.

## Como criar seus próprios testes

Um teste é composto de um `input` (o texto que será fornecido para o programa) e um `output` (o texto que o programa deve retornar para esse input) e pode ter opcionalmente um `label` para facilitar a identificação do teste.

Seus casos de teste personalizados podem ser escritos diretamente aqui na descrição do problema dentro de um fenced code block com a linguagem `toml` ou em um arquivo `tests.toml` na pasta do problema. O TKO irá carregar automaticamente os testes quando a tarefa for aberta ou executada novamente.

Exemplo de teste para ler dois números, um por linha, e imprimir a soma e a subtração deles.

Se quiser habilitar esses casos de teste e ver funcionando, insira algo no input e no output.

```toml
# Exemplo de entrada em uma linha
[[tests]]
input = ''
output = ''

# Exemplo de entrada em múltiplas linhas
# [[tests]]
# input = '''
# 1
# 2
# '''
# output = '''
# 3
# 4
# '''
```

