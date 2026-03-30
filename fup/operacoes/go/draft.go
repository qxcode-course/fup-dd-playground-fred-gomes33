package main

import "fmt"

func main() {
	var n1, n2, som, sub, mult, div, rest int
    var divf float64
	fmt.Scan(&n1, &n2, &som, &sub, &mult, &div, &rest, &divf)
	som = n1 + n2
	sub = n1 - n2
	mult = n1 * n2
	div = n1 / n2
	rest = n1 % n2
    divf = float64(n1) / float64(n2)
	fmt.Printf("%d\n%d\n%d\n%.2f\n%d\n", som, sub, mult, divf, rest)
}
