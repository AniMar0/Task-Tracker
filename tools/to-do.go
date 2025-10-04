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
	file, err := os.Create("db.json")
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	return nil
}

func GetTodos() ([]Todo, error) {
	var Todos []Todo
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Todos); err != nil {
		if err.Error() == "EOF" {
			return []Todo{}, nil
		}
		return nil, err
	}
	return Todos, nil
}

func AddTodo(todo Todo) error {
	var (
		err   error
		Todos []Todo
	)
	Todos, err = GetTodos()
	if err != nil {
		return err
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
	file, err := os.Open("db.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(Todos); err != nil {
		return err
	}
	return nil
}
