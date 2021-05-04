package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	charset := "abcdefghijklmnoprstuvyzxABCDEFGHIJKLMNOPRSTUVYZX1234567890!*-_~,%&/"
	lenght := 15
	byteArray := make([]byte, lenght)
	for i := range byteArray {
		byteArray[i] = charset[source.Int63()%int64(len(charset))]
	}
	pass := string(byteArray)
	fmt.Println(pass)
}
