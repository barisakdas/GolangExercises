package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=997d9be77b8cfe25ffc8808b509bd530&lat=41.058316&lon=29.004077")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	//result := fmt.Sprintf(bodyString)
	//return result
}
