package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FormatingToDo(todo Todo) string {
	return fmt.Sprintf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n", todo.ID, todo.Description, todo.Status, FormatingTime(todo.CreatedAt), FormatingTime(todo.UpdatedAt))
}

func FormatingAddTask(id int) string {
	return fmt.Sprintf("Task added successfully (ID: %d)\n", id)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func FormatingTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatingUpdateTask(id int) string {
	return fmt.Sprintf("Task updated successfully (ID: %d)\n", id)
}

func FormatingInput(input []string) []string {
	if len(input) <= 1 {
		return input
	}

	switch input[0] {
	case "add":
		return []string{input[0], strings.Join(input[1:], " ")[1 : len(strings.Join(input[1:], " "))-1]}
	case "update":
		if strings.Contains(input[1], "-") {
			return input
		}
		return []string{input[0], input[1], strings.Join(input[2:], " ")[1 : len(strings.Join(input[2:], " "))-1]}
	}

	return input
}
