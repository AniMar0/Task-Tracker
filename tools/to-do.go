package tools

import (
	"encoding/json"
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
