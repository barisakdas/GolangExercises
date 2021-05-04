/* UPDATED VERSION */

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Option struct { // Dışarıdan son kullanıcı şifre uzunluğunu ve şifre setinde nelerin olacağını belirleyecek.
	lenght            int
	upperCase         bool
	lowerCase         bool
	numbers           bool
	specialCharacters bool
}

func main() {
	opt := Option{
		lenght:            15,
		lowerCase:         true,
		upperCase:         true,
		numbers:           false,
		specialCharacters: false,
	}
	pass, err := GeneratePassword(opt)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Your password: ", pass)
	}
}

func GeneratePassword(opt Option) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	charLowerCase := "abcdefghijklmnoprstuvyzx"
	charUpperCase := "ABCDEFGHIJKLMNOPRSTUVYZX"
	charNumbers := "1234567890"
	charSpecialCharacters := "!*-_~,%&/+[]"

	byteArray := make([]byte, opt.lenght)
	charset := ""

	if opt.lowerCase {
		charset += charLowerCase
	}
	if opt.upperCase {
		charset += charUpperCase
	}
	if opt.numbers {
		charset += charNumbers
	}
	if opt.specialCharacters {
		charset += charSpecialCharacters
	}

	if len(charset) < 1 {
		return "NON-GENERATED", errors.New("Did you forget choose any charset?")
	}

	for i := range byteArray {
		byteArray[i] = charset[source.Int63()%int64(len(charset))]
	}
	pass := string(byteArray)

	return pass, nil
}
