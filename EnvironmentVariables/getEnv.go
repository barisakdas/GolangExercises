package main

import (
	"fmt"
	"os"
)

func main() {
	// İstediğiniz herhangi bir ortam değişkenini almak için bu kullanım mevcut.
	userName := os.Getenv("USERNAME")
	fmt.Println("USERNAME: ", userName)

	userDomain := os.Getenv("USERDOMAIN")
	fmt.Println("USERDOMAIN: ", userDomain)

	goroot := os.Getenv("GOROOT")
	fmt.Println("GOROOT: ", goroot)

	gopath := os.Getenv("GOPATH")
	fmt.Println("GOPATH: ", gopath)

	homepath := os.Getenv("HOMEPATH")
	fmt.Println("HOMEPATH: ", homepath)
}
