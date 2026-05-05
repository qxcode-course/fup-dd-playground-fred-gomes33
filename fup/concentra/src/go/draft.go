package main
import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	fmt.Print("[ ")
	i := a
	j := b
	for i <= b {
		fmt.Printf("%d %d ", i, j)
		i++
		j--
	}

	fmt.Println("]")
}