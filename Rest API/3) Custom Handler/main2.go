package main

import (
	"fmt"
	"net/http"
)

func main() {

	s1 := &Student{"Barış Akdaş"}

	mux := http.NewServeMux()      // Her seferinde http kütüphanesini kullanarak http.HandlFunc yazmamamız için kullanılır.
	mux.Handle("/studentName", s1) // Nesneyi direkt kullanabilmen için `SurveHTTP` metoduna implemente edilmesi gerekiyor.
	http.ListenAndServe(":8000", mux)
}

type Student struct { // Bu nesne bizim custom handler nesnemiz
	FirstName string
}

func (std Student) ServeHTTP(res http.ResponseWriter, r *http.Request) { // Fonksiyonun ismi `ServeHTTP` olmazsa yukarda mux.Handle içinde direkt nesneyi kullanamıyoruz.
	//io.WriteString(res, "custom handler metodu çalıştı") // res.Write metodu ile aynı mantıkta çalışacak.
	fmt.Fprintf(res, "Student Name : "+std.FirstName)
}

// ## Test Url : `http://localhost:8000/studentName`
