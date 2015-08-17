package main

import "github.com/gorilla/mux"
import "net/http"
import "html/template"

var templates = template.Must(template.ParseFiles("index.html"))

type Page struct {
}

func RunWeb() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)

	err := http.ListenAndServe("localhost:8899", router)
	check(err)
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	templates.ExecuteTemplate(response, "index.html", Page{})
}
