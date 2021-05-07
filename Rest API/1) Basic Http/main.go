package main

import (
	"log"
	"net/http"
)

// url tetiklenince çalışmasını istediğimiz fonksiyon.
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bu alan about sayfasının url i tetiklendiğinde çalışacak olan fonksiyonun içi."))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Dünya!")) // ResponseWriter nesnesi bir byte dizisini yazdırır.

func main() {
	// Bu `localhost:8000/` url ine gelen istek ikinci parametre de alınan fonksiyonu çalıştırır.
	http.HandleFunc("/", indexHandler)

	// Bu `localhost:8000/about` url ine gelen istek ikinci parametre de alınan fonksiyonu çalıştırır.
	http.HandleFunc("/about", aboutHandler)

	// Nerede yayınlayacağımızı burada söylüyoruz. `ListenAndServeTLS` fonkdiyonu https yayını için kullanılıyor.
	http.ListenAndServe(":8000", nil)
}

// ## Test Url : `http://localhost:8000/`
// ## Test Url : `http://localhost:8000/about`
