package tools

import (
	"fmt"
	"strconv"
)

func FormatingToDo(todo Todo) string {
	return fmt.Sprintf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n", todo.ID, todo.Description, todo.Status, todo.CreatedAt, todo.UpdatedAt)
}

func FormatingAddTask(id int) string {
	return fmt.Sprintf("Task added successfully (ID: %d)\n", id)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
