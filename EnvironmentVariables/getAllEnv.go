package main

import (
	"fmt"
	"os"
)

func main() {
	// Tüm ortam değişkenlerini verecektir.
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
