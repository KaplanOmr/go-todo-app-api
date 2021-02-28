package database

import (
	"database/sql"
	"fmt"

	. "../helpers"
	. "../models"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

func conn() *sql.DB {

	var conf DatabaseConfig
	_, err := toml.DecodeFile("../database/config.toml", &conf)
	CheckErr(err)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)

	db, err := sql.Open("mysql", connString)

	return db
}

func All() TodoList {

	var respond TodoList

	db := conn()
	defer db.Close()

	result, err := db.Query("SELECT * FROM todo")
	CheckErr(err)

	for result.Next() {
		var todo Todo

		err := result.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Date)

		CheckErr(err)

		respond.Todo = append(respond.Todo, todo)
	}

	return respond
}

func Get(id int) Todo {

	var todo Todo

	db := conn()
	defer db.Close()

	err := db.QueryRow("SELECT * FROM todo WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Date)
	CheckErr(err)

	return todo
}

func GetDate(date string) TodoDateMappings {

	var todoDate TodoDateMappings
	var todo Todo

	todoDate.Date = date

	unixDate := ToUnix(date)

	db := conn()
	defer db.Close()
	result, err := db.Query("SELECT * FROM todo WHERE date = ?", unixDate)
	CheckErr(err)

	for result.Next() {
		err = result.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Date)
		CheckErr(err)

		todoDate.Todo = append(todoDate.Todo, todo)
	}

	return todoDate
}

func Add(newTodo NewTodo) bool {

	date := CurrDate()

	db := conn()
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO todo SET title = ?, description = ?, status = ?, date = ?",
		newTodo.Title,
		newTodo.Description,
		newTodo.Status,
		date)
	defer insert.Close()

	if err != nil {
		return false
	}

	return true

}

func Up(todo Todo) bool {

	db := conn()
	defer db.Close()

	update, err := db.Query(
		"UPDATE todo SET title = ?, description = ?, status = ? WHERE id = ?",
		todo.Title,
		todo.Description,
		todo.Status,
		todo.ID)
	defer update.Close()

	if err != nil {
		return false
	}

	return true

}

func Delete(id int) bool {
	db := conn()
	defer db.Close()

	update, err := db.Query(
		"DELETE FROM todo WHERE id = ?",
		id)
	defer update.Close()

	if err != nil {
		return false
	}

	return true
}
