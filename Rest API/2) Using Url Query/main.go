package main

import (
	"log"
	"net/http"
)

func main() {
	// Bu `localhost:8080/` url ine gelen istek ikinci parametre de alınan fonksiyonu çalıştırır.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)

	// Nerede yayınlayacağımızı burada söylüyoruz. `ListenAndServeTLS` fonkdiyonu https yayını için kullanılıyor.
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"] // `.URL.Query()["key"]` url üzerindeki ? den sonra gelen query leri yakalamak için kullanılır. Örnek url: `localhost:8000/?key=value`

	if !ok || len(keys[0]) < 1 {
		w.Write([]byte("Url Param 'key' is missing"))
		log.Println("Url Param 'key' is missing") // Terminal üzerindeki konsol ekranına log atar.
		return
	}
	key := keys[0]

	w.Write([]byte(string(key)))
	log.Println("Url Param 'key' is: " + string(key))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	number := r.FormValue("number")

	data := "## Girilen url: " + "localhost:8080/search?name=" + name + "&age=" + age + "&number=" + number
}

// ## Test Url : `http://localhost:8080/`
// ## Test Url : `http://localhost:8080/search?name=baris&age=28&number=1453`
