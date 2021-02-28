package models

type Todo struct {
	ID          int    `json:"id"`
	Date        int    `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TodoDateMappings struct {
	Date string `json:"date"`
	Todo []Todo `json:"todo"`
}

type TodoList struct {
	Todo []Todo `json:"todo"`
}
