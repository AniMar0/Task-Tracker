package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func NormalizeSpaces(s string) string {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

func FormatingToDo(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	maxIDLen := len("ID")
	maxDescLen := len("Description")
	maxStatusLen := len("Status")
	maxCreatedLen := len("Created At")
	maxUpdatedLen := len("Updated At")

	for _, todo := range todos {
		desc := NormalizeSpaces(todo.Description)
		status := NormalizeSpaces(todo.Status)
		created := NormalizeSpaces(FormatingTime(todo.CreatedAt))
		updated := NormalizeSpaces(FormatingTime(todo.UpdatedAt))

		if l := len(fmt.Sprintf("%d", todo.ID)); l > maxIDLen {
			maxIDLen = l
		}
		if l := len(desc); l > maxDescLen {
			maxDescLen = l
		}
		if l := len(status); l > maxStatusLen {
			maxStatusLen = l
		}
		if l := len(created); l > maxCreatedLen {
			maxCreatedLen = l
		}
		if l := len(updated); l > maxUpdatedLen {
			maxUpdatedLen = l
		}
	}

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
			maxDescLen, NormalizeSpaces(todo.Description),
			maxStatusLen, NormalizeSpaces(todo.Status),
			maxCreatedLen, NormalizeSpaces(FormatingTime(todo.CreatedAt)),
			maxUpdatedLen, NormalizeSpaces(FormatingTime(todo.UpdatedAt)),
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

func HelpMenu() {
	commands := []struct {
		cmd  string
		desc string
	}{
		{"list", "Lists all tasks"},
		{"list [todo|in-progress|done]", "Lists tasks of a specific status"},
		{"mark-[todo|in-progress|done] [ID]", "Marks a task with the given status"},
		{`add ["Description"]`, "Adds a new task"},
		{`update [ID] ["Description"]`, "Updates a task description"},
		{"delete [ID]", "Deletes a task"},
		{"exit", "Exits the program"},
	}

	maxCmdLen := len("Command")
	maxDescLen := len("Description")
	for _, c := range commands {
		if l := len(c.cmd); l > maxCmdLen {
			maxCmdLen = l
		}
		if l := len(c.desc); l > maxDescLen {
			maxDescLen = l
		}
	}

	totalWidth := maxCmdLen + maxDescLen + 5 // 5 = for | and spaces

	printLine := func() {
		fmt.Printf("+-%s-+-%s-+\n",
			strings.Repeat("-", maxCmdLen),
			strings.Repeat("-", maxDescLen),
		)
	}

	// Print centered title
	title := "Task Tracker CLI - Help Menu"
	padding := ((totalWidth - len(title)) / 2)
	if padding < 0 {
		padding = 0
	}
	fmt.Println(strings.Repeat("=", totalWidth))
	fmt.Printf("%s%s\n", strings.Repeat(" ", padding), title)
	fmt.Println(strings.Repeat("=", totalWidth))

	// Table header
	printLine()
	fmt.Printf("| %-*s | %-*s |\n", maxCmdLen, "Command", maxDescLen, "Description")
	printLine()

	// Table rows
	for _, c := range commands {
		fmt.Printf("| %-*s | %-*s |\n", maxCmdLen, c.cmd, maxDescLen, c.desc)
	}

	printLine()
}
