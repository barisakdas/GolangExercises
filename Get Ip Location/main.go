package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	ip := "8.8.8.8"
	cmd := "https://api.ipgeolocation.io/ipgeo?ip=" + ip + "&apiKey=03d1f217ffff4394821651e0b8f4884a"

	resp, err := http.Get(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bodyString)
}
