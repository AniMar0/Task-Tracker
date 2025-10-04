package tools

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Todo struct {
	ID          int       `json:"id,omitempty"`
	Description string    `json:"description"`
	Status      string    `json:"status,omitempty"` //(todo, in-progress, done)
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func MakeToDoFile() error {
	file, err := os.OpenFile("db.json", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	return nil
}

func GetTodos(GetType string) ([]Todo, error) {
	var Todos []Todo
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Todos); err != nil {
		return nil, err
	}
	if GetType == "list" {
		return Todos, nil
	}
	var TodoSplit []Todo
	for _, todo := range Todos {
		switch GetType {
		case "todo":
			if todo.Status == "todo" {
				TodoSplit = append(TodoSplit, todo)
			}
		case "in-progress":
			if todo.Status == "in-progress" {
				TodoSplit = append(TodoSplit, todo)
			}
		case "done":
			if todo.Status == "done" {
				TodoSplit = append(TodoSplit, todo)
			}
		default:
			return nil, errors.New("invalid GetType")
		}

	}
	return TodoSplit, nil
}

func AddTodo(todo Todo) error {
	var (
		err   error
		Todos []Todo
	)
	Todos, err = GetTodos("list")
	if err != nil {
		if err.Error() != "EOF" {
			return err

		} else {
			Todos = []Todo{}
		}
	}
	for _, t := range Todos {
		if t.ID == todo.ID {
			return errors.New("ID already exists")
		}
	}
	Todos = append(Todos, todo)
	SaveTodos(Todos)
	return nil
}

func SaveTodos(Todos []Todo) error {
	file, err := os.Create("db.json")
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(Todos); err != nil {
		return err
	}
	return nil
}

func DeleteTodo(idInt int, Todos []Todo) error {
	for i, todo := range Todos {
		if todo.ID == idInt {
			Todos = append(Todos[:i], Todos[i+1:]...)
			SaveTodos(Todos)
			return nil
		}
	}
	return errors.New("ID not found")
}

func GetTodo(id int, Todos []Todo) Todo {
	for _, todo := range Todos {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}

func UpdateTodo(idInt int, Update string, Todos []Todo) error {
	for i, todo := range Todos {
		if todo.ID == idInt {
			switch Update {
			case "mark-in-progress":
				Todos[i].Status = "in-progress"
			case "mark-done":
				Todos[i].Status = "done"
			case "mark-todo":
				Todos[i].Status = "todo"
			default:
				Todos[i].Description = Update
			}
			Todos[i].UpdatedAt = time.Now()
			SaveTodos(Todos)
			return nil
		}
	}
	return errors.New("ID not found")
}
