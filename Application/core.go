package main

import (
	"log"
	"net/http"
)

const baseURL = "http://bari.opendata.planetek.it"

// FetchData gets data from the external data source (REST service)
func FetchData(endpoint string) string{
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		log.Println("[HTTP_FETCH] Error:", err.Error())
	}
	var data []byte
	resp.Body.Read(data) //TODO: error handling
	str := string(data)
	return str
}