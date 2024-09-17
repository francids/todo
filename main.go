package main

import (
	"fmt"
	"os"
	"strings"
	todo "todo/cmd"
)

func main() {
	todo.InitDB("todos.sqlite3")

	if len(os.Args) < 2 {
		fmt.Println("Uso: todo <comando> [argumentos]")
		return
	}

	command := os.Args[1]
	switch command {
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Uso: todo new <tarea>")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		todo.AddTodo(task)
	case "list":
		todo.ListTodos()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Uso: todo remove <id>")
			return
		}
		id := os.Args[2]
		todo.RemoveTodoByID(id)
	default:
		fmt.Println("Comando no reconocido.")
	}
}
