package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FormatingToDo(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// نحسب الطول الأعظمي لكل عمود
	maxIDLen := len("ID")
	maxDescLen := len("Description")
	maxStatusLen := len("Status")
	maxCreatedLen := len("Created At")
	maxUpdatedLen := len("Updated At")

	for _, todo := range todos {
		if l := len(fmt.Sprintf("%d", todo.ID)); l > maxIDLen {
			maxIDLen = l
		}
		if l := len(todo.Description); l > maxDescLen {
			maxDescLen = l
		}
		if l := len(todo.Status); l > maxStatusLen {
			maxStatusLen = l
		}
		if l := len(FormatingTime(todo.CreatedAt)); l > maxCreatedLen {
			maxCreatedLen = l
		}
		if l := len(FormatingTime(todo.UpdatedAt)); l > maxUpdatedLen {
			maxUpdatedLen = l
		}
	}

	// function صغيرة باش نرسم سطر
	printLine := func() {
		fmt.Printf("+-%s-+-%s-+-%s-+-%s-+-%s-+\n",
			strings.Repeat("-", maxIDLen),
			strings.Repeat("-", maxDescLen),
			strings.Repeat("-", maxStatusLen),
			strings.Repeat("-", maxCreatedLen),
			strings.Repeat("-", maxUpdatedLen),
		)
	}

	// Header
	printLine()
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		maxIDLen, "ID",
		maxDescLen, "Description",
		maxStatusLen, "Status",
		maxCreatedLen, "Created At",
		maxUpdatedLen, "Updated At",
	)
	printLine()

	// Rows
	for _, todo := range todos {
		fmt.Printf("| %-*d | %-*s | %-*s | %-*s | %-*s |\n",
			maxIDLen, todo.ID,
			maxDescLen, todo.Description,
			maxStatusLen, todo.Status,
			maxCreatedLen, FormatingTime(todo.CreatedAt),
			maxUpdatedLen, FormatingTime(todo.UpdatedAt),
		)
	}

	// Footer
	printLine()
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
