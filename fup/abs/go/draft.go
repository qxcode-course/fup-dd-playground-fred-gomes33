package main

import "fmt"

func main() {
	var num1, num2, dif int
	fmt.Scan(&num1, &num2, &dif)

	if num1 < 0 && num2 < 0 {
		num1, num2 = num1*-1, num2*-1
	} else {
		num1, num2 = num1*1, num2*1
	}

	dif = num2 - num1

	if dif < 0 {
		dif = dif * -1
	} else {
		dif = dif * 1
	}

	fmt.Println(dif)
}
