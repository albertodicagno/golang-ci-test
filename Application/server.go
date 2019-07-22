package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// StartWebService starts the web server
func StartWebService() {
	http.HandleFunc("/api/v1/Linee", GetLines)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}

func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {

	tmpl = fmt.Sprintf("views/%s", tmpl)
	t, err := template.ParseFiles(tmpl)

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

// Home handles home page rendering
func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Message:  "Success!",
		Language: "Go Lang",
	}
	render(w, "index.html", pageVars)
}
