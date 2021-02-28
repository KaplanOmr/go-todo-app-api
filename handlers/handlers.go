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
	var resNot RespondNotice
	err := json.NewDecoder(r.Body).Decode(&todo)
	CheckErr(err)

	upCheck := database.Up(todo)

	if upCheck {
		resNot.Status = "success"
		resNot.Notice = "Todo updated"
	} else {
		resNot.Status = "success"
		resNot.Notice = "Todo not updated"
	}

	respond, _ := json.Marshal(resNot)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo NewTodo
	var resNot RespondNotice
	err := json.NewDecoder(r.Body).Decode(&todo)
	CheckErr(err)

	addCheck := database.Add(todo)

	if addCheck {
		resNot.Status = "success"
		resNot.Notice = "Todo added"
	} else {
		resNot.Status = "error"
		resNot.Notice = "Todo not added"
	}

	respond, _ := json.Marshal(resNot)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(respond))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	var resNot RespondNotice
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	deleteCheck := database.Delete(key)

	if deleteCheck {
		resNot.Status = "success"
		resNot.Notice = "Todo deleted"
	} else {
		resNot.Status = "error"
		resNot.Notice = "Todo not deleted"
	}

	respond, _ := json.Marshal(resNot)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respond))
}
