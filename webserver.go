package main

import "github.com/gorilla/mux"
import "net/http"
import "fmt"

func RunWeb() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	err := http.ListenAndServe("localhost:8899", router)
	check(err)
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	fmt.Fprintf(response, "Hey there!\n")
}
