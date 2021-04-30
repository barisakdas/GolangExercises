package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("##############################################################################")
	fmt.Println("\t\tÇarpım tablosu alıştırmalarına hoş geldiniz.")
	fmt.Println("##############################################################################")
	level := getLevel()
	if level == 1 {
		getQuestion(1, 5)
	} else if level == 2 {
		getQuestion(5, 10)
	} else if level == 3 {
		getQuestion(10, 25)
	} else if level == 4 {
		getQuestion(25, 100)
	} else {
		fmt.Println("Hatalı bir seçim yaptınız!")
		getLevel()
	}
}

func getLevel() int {
	var level string
	fmt.Println("Zorluk Seviyesi Seçiniz: (Çıkış=-1)")
	fmt.Println("Kolay : 1")
	fmt.Println("Orta : 2")
	fmt.Println("Zor : 3")
	fmt.Println("Çok Zor : 4")
	fmt.Scanf("%s", &level)
	intLevel, _ := strconv.Atoi(level)
	return intLevel
}

func getQuestion(range1 int, range2 int) {

	randomNum1 := rand.Intn(range2-range1) + 1 + range1
	randomNum2 := rand.Intn(range2-range1) + 1 + range1

	fmt.Printf("\t %d x %d kaça eşittir?\n", randomNum1, randomNum2)
	response := 0
	fmt.Print("Sonuç >> ")
	fmt.Scanf("%d", &response)
	result := randomNum1 * randomNum2
	if response == result {
		fmt.Printf("\t#### Doğru. Tebrikler!! ####\n")
		fmt.Printf("_______________________________________________________ \n")
	} else {
		fmt.Printf("\t#### Yanlış. Cevap %d olacaktı!! ####\n", randomNum1*randomNum2)
		fmt.Printf("_______________________________________________________ \n")
	}

}
