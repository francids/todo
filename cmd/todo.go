package todo

import (
	"fmt"
	"todo/utils"
)

type Todo struct {
	ID   string
	Task string
}

var todos []Todo

func AddTodo(task string) {
	id := utils.GenerateID(task)
	todo := Todo{ID: id, Task: task}
	todos = append(todos, todo)
	fmt.Printf("Tarea '%s' creada con ID %s\n", task, todo.ID)
}

func ListTodos() {
	if len(todos) == 0 {
		fmt.Println("No hay tareas.")
		return
	}
	for _, todo := range todos {
		fmt.Printf("%s: %s\n", todo.ID, todo.Task)
	}
}
