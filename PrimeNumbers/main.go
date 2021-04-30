package main

import "fmt"

func main() {
	var primeNumbers []int
	var number int
	fmt.Print("How many prime numbers are you looking for? :")
	fmt.Scanf("%d", &number)

	for i := 2; i < number; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				//break
			}
		}
		if flag {
			primeNumbers = append(primeNumbers, i)
		}
	}
	fmt.Println(primeNumbers)
}
