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

func RemoveTodoByID(id string) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Printf("Tarea '%s' eliminada\n", todo.Task)
			return
		}
	}
	fmt.Printf("No se encontró tarea con ID %s\n", id)
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
