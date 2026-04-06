package main

import "fmt"

func floor(num float64) int {
	inteiro := int(num)
	if num >= 0 {
		return inteiro
	} else {
		if num == float64(inteiro) {
			return inteiro
		} else {
			return inteiro - 1
		}
	}
}

func ceil(num float64) int {
	inteiro := int(num)
	if num >= 0 {
		if num == float64(inteiro) {
			return inteiro
		} else {
			return inteiro + 1
		}
	} else {
		return inteiro
	}
}

func round(num float64) int {
	inteiro := int(num)
	frac := num - float64(inteiro)
	if num >= 0 {
		if frac >= 0.5 {
			return inteiro + 1
		} else {
			return inteiro
		}
	} else {
		if frac <= -0.5 {
			return inteiro - 1
		} else {
			return inteiro
		}
	}
}

func main() {
	var op string
	var num float64
	fmt.Scan(&op, &num)
	switch op[0] {
	case 'f':
		fmt.Println(floor(num))
	case 'c':
		fmt.Println(ceil(num))
	case 'r':
		fmt.Println(round(num))
	}
}
