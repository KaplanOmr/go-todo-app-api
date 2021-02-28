package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "../database"
	. "../helpers"
	. "../models"
	"github.com/gorilla/mux"
)

func AllTodo(w http.ResponseWriter, r *http.Request) {

	respond, err := json.Marshal(database.All())

	CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	respond, err := json.Marshal(database.Get(key))

	CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}

func GetDateTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, _ := vars["date"]

	respond, err := json.Marshal(database.GetDate(key))

	CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}

func UpTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var respond string
	err := json.NewDecoder(r.Body).Decode(&todo)
	CheckErr(err)

	upCheck := database.Up(todo)

	if upCheck {
		respond = "success"
	} else {
		respond = "error"
	}

	CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo NewTodo
	var respond string
	err := json.NewDecoder(r.Body).Decode(&todo)
	CheckErr(err)

	upCheck := database.Add(todo)

	if upCheck {
		respond = "success"
	} else {
		respond = "error"
	}

	CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(respond))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	var respond string
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	deleteCheck := database.Delete(key)

	if deleteCheck {
		respond = "success"
	} else {
		respond = "error"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}
