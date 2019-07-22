package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	baseURL     string = "http://bari.opendata.planetek.it"
	getLineeURI string = "/OrariBus/v2.1/OpenDataService.svc/REST/rete/Linee"
)

// FetchData gets data from the external data source (REST service)
func FetchData(endpoint string) []byte {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		log.Println("[HTTP_FETCH] Error:", err.Error())
	}
	var data = make([]byte, resp.ContentLength)
	resp.Body.Read(data) //TODO: error handling
	log.Printf("[HTTP_FETCH] fetched %d bytes from %s\n", len(data), endpoint)
	return data
}

func GetLines(w http.ResponseWriter, req *http.Request) {
	d := FetchData(getLineeURI)
	var linee []Linea
	json.Unmarshal(d, linee)
	w.Write(d)
}
