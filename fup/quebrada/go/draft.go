package main

import "fmt"

func main() {
	var n1, n2, div, rest, res int
	var rest2 float64
	fmt.Scan(&n1, &n2, &div, &rest, &res, &rest2)
	div = n1 / n2
	rest = div - 2
	res = n1 % n2
	rest2 = float64(n1) / float64(n2)
	fmt.Println(div)
	fmt.Println(res)
	fmt.Printf("%.2f\n", rest2)
}
