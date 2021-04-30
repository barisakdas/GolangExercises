package main

import "fmt"

func main() {
	fmt.Println("Faktorieli hesaplanacak sayıyı giriniz.")
	var number int
	fmt.Scanf("%d", &number)
	result := CalculateFactorial(number)
	fmt.Println("Faktoriel: ", result)
}

func CalculateFactorial(number int) int {
	var result int
	if number > 1 {
		result = number * CalculateFactorial(number-1)
	} else {
		return 1
	}
	return result
}
