package main

import (
	"net/http"

	. "../handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/all", AllTodo).Methods("GET")
	r.HandleFunc("/get/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/getdate/{date}", GetDateTodo).Methods("GET")
	r.HandleFunc("/add", AddTodo).Methods("POST")
	r.HandleFunc("/up", UpTodo).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteTodo).Methods("DELETE")

	serve := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	serve.ListenAndServe()
}
