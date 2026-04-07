package main

import "fmt"
import "math"

func main() {
	var a, b, c, del, x1, x2 float64
	fmt.Scan(&a, &b, &c, &del, &x1, &x2)

	del = b*b - 4*a*c

	x1 = (-b + math.Sqrt(del)) / (2 * a)
	x2 = (-b - math.Sqrt(del)) / (2 * a)

	if del > 0 {
		fmt.Printf("%.2f\n", x1)
		fmt.Printf("%.2f\n", x2)
	} else if del == 0 {
		fmt.Printf("%.2f\n", x1)
	} else if del < 0 {
		fmt.Println("nao ha raiz real")
	}
}
