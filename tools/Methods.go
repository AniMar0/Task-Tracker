package tools

import (
	"errors"
	"fmt"
	"time"
)

type Request struct {
	Method      string
	List        string
	Description string
	ID          int
	Status      string
}

func TodoLC(r Request) error {
	switch r.Method {
	case "list":
		todos, err := GetTodos(r.List)
		if err != nil {
			return err
		}
		for _, todo := range todos {
			fmt.Println(FormatingToDo(todo))
		}
	case "add":
		Todos, err := GetTodos("list")
		if err != nil {
			return err
		}
		var todo Todo
		todo.Description = r.Description
		todo.Status = "todo"
		todo.CreatedAt = time.Now()
		lastID := Todos[len(Todos)-1].ID
		todo.ID = lastID + 1

		err = AddTodo(todo)
		if err != nil {
			return err
		}
		fmt.Println(FormatingAddTask(todo.ID))
	case "update":
		Todos, err := GetTodos("list")
		if err != nil {
			return err
		}
		if r.Description == "" {
			r.Description = r.Status
		}
		err = UpdateTodo(r.ID, r.Description, Todos)
		if err != nil {
			return err
		}
	case "delete":
	default:
		return errors.New("method not found")
	}
	return nil
}
