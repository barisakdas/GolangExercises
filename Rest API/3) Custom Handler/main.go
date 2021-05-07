package main

import (
	"io"
	"net/http"
)

func main() {

	var bmw car

	mux := http.NewServeMux()         // Her seferinde http kütüphanesini kullanarak http.HandlFunc yazmamamız için kullanılır.
	mux.Handle("/customHandler", bmw) // Nesneyi direkt kullanabilmen için `SurveHTTP` metoduna implemente edilmesi gerekiyor.
	http.ListenAndServe(":8000", mux)
}

type car int

func (c car) ServeHTTP(res http.ResponseWriter, r *http.Request) { // Fonksiyonun ismi `ServeHTTP` olmazsa yukarda mux.Handle içinde direkt nesneyi kullanamıyoruz.
	io.WriteString(res, "custom handler metodu çalıştı") // res.Write metodu ile aynı mantıkta çalışacak.
}

// ## Test Url : `http://localhost:8000/customHandler`
