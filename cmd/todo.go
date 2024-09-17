package todo

import (
	"fmt"
	"log"
	"todo/utils"
)

type Todo struct {
	ID   string
	Task string
}

func AddTodo(task string) {
	id := utils.GenerateID(task)
	todo := Todo{ID: id, Task: task}

	insertTodoSQL := `INSERT INTO todos(id, task) VALUES (?, ?)`
	statement, err := db.Prepare(insertTodoSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(todo.ID, todo.Task)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tarea '%s' creada con ID %s\n", task, todo.ID)
}

func RemoveTodoByID(id string) {
	deleteTodoSQL := `DELETE FROM todos WHERE id = ?`
	statement, err := db.Prepare(deleteTodoSQL)
	if err != nil {
		log.Fatal(err)
	}
	result, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Printf("No se encontró tarea con ID %s\n", id)
	} else {
		fmt.Printf("Tarea con ID %s eliminada\n", id)
	}
}

func ListTodos() {
	row, err := db.Query("SELECT id, task FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id, task string
		row.Scan(&id, &task)
		fmt.Printf("%s: %s\n", id, task)
	}
}
