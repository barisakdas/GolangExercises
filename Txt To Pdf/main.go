package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	fileName := "test.txt"

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("test.pdf")

	fmt.Println("PDF Created.")

}
