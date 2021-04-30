package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// please try carefully
	CreateFile()
	// GetDataFromFile()
	// GetDataFromFile()
	// WriteData()
	// SelectFile()
	// GetFileMetaData()
	// DeleteFile()
}

func CreateFile() {

	file, err := os.Create("notes.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func SelectFile() {
	file, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

func GetDataFromFile() {
	file, err := os.OpenFile("notes.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetFileMetaData() {
	file, err := os.Stat("notes.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name: ", file.Name())
	fmt.Println("File Size: ", file.Size())
	fmt.Println("File Permission	: ", file.Mode())
	fmt.Println("File Created Date: ", file.ModTime())
	fmt.Println("File Is Directory: ", file.IsDir())
	fmt.Println("File Source: ", file.Sys())
}

func WriteData() {
	file, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	len, err := file.WriteString("Hello, this string will be added to the end of the file")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Byte size written: " + strconv.Itoa(len))
	}
}

func DeleteFile() {
	err := os.Remove("notes.txt")
	if err != nil {
		log.Fatal(err)
	}

	err2 := os.RemoveAll("/home/user/test")
	if err2 != nil {
		log.Fatal(err2)
	}
}
