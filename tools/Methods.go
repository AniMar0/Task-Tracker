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
			if err.Error() == "EOF" {
				return errors.New("no tasks found")
			}
			return err
		}
		FormatingToDo(todos)
	case "add":
		Todos, err := GetTodos("list")
		if err != nil {
			if err.Error() != "EOF" {
				return err

			} else {
				Todos = []Todo{}
			}
		}
		var todo Todo
		todo.Description = r.Description
		todo.Status = "todo"
		todo.CreatedAt = time.Now()
		todo.UpdatedAt = time.Now()
		if len(Todos) == 0 {
			todo.ID = 1
		} else {
			lastID := Todos[len(Todos)-1].ID
			todo.ID = lastID + 1
		}

		err = AddTodo(todo)
		if err != nil {
			return err
		}
		fmt.Println(FormatingAddTask(todo.ID))
	case "update":
		Todos, err := GetTodos("list")
		if err != nil {
			if err.Error() != "EOF" {
				return err

			} else {
				Todos = []Todo{}
			}
		}
		if r.Description == "" {
			r.Description = r.Status
		}
		err = UpdateTodo(r.ID, r.Description, Todos)
		if err != nil {
			return err
		}
	case "delete":
		Todos, err := GetTodos("list")
		if err != nil {
			if err.Error() != "EOF" {
				return err

			} else {
				Todos = []Todo{}
			}
		}
		if r.Description == "" {
			r.Description = r.Status
		}
		err = DeleteTodo(r.ID, Todos)
		if err != nil {
			return err
		}
	default:
		return errors.New("method not found")
	}
	return nil
}
