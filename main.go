package main

import (
	"Task-Tracker/tools"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tools.MakeToDoFile()
	reader := bufio.NewReader(os.Stdin)
	counter := 0
	for {
		if counter == 0 {
			fmt.Println("Please provide a command or type \"help\" for more information")
			counter++
		}

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		Args := tools.FormatingInput(strings.Fields(input))

		if len(Args) == 0 {
			fmt.Println("Please provide a command or type \"help\" for more information")
			continue
		}
		if Args[0] == "exit" {
			break
		} else if Args[0] == "help" {
			fmt.Print("=========================================================================\n",
				"Commands: \n",
				"=========================================================================\n",
				"list - Lists all tasks \n",
				"list [todo|in-progress|done] - Lists all tasks of a specific status \n",
				"mark-[in-progress|done|todo] [ID] - Marks a task as [in-progress|done|todo] \n",
				"add [\"Description\"] - Adds a task \n",
				"update [ID] [\"Description\"]- Updates a task \n",
				"delete [ID] - Deletes a task \n",
				"exit - Exits the program \n",
				"=========================================================================\n",
			)
		} else if Args[0] == "list" {
			var Request tools.Request
			Request.Method = "list"
			if len(Args) == 2 {
				Request.List = Args[1]
			} else {
				Request.List = "list"
			}
			err := tools.TodoLC(Request)
			if err != nil {
				fmt.Println(err)
				counter = 0
			}
		} else if Args[0] == "add" {
			var Request tools.Request
			fmt.Println(Args)
			if len(Args) == 2 {
				Request.Method = Args[0]
				Request.Description = Args[1]
				err := tools.TodoLC(Request)
				if err != nil {
					fmt.Println(err)
					counter = 0
				}
			} else {
				fmt.Println("Please provide a description")
			}
		} else if Args[0] == "update" {
			var Request tools.Request
			if len(Args) == 3 {
				Request.Method = Args[0]
				Request.ID = tools.StringToInt(Args[1])
				Request.Description = Args[2]
				err := tools.TodoLC(Request)
				if err != nil {
					fmt.Println(err)
					counter = 0
				}
			} else {
				fmt.Println("Please provide an ID and a description")
			}
		} else if Args[0] == "delete" {
			var Request tools.Request
			if len(Args) == 2 {
				Request.Method = Args[0]
				Request.ID = tools.StringToInt(Args[1])
				err := tools.TodoLC(Request)
				if err != nil {
					fmt.Println(err)
					counter = 0
				}
			} else {
				fmt.Println("Please provide an ID")
			}

		} else if len(Args) == 2 && strings.HasPrefix(Args[0], "mark-") {
			var Request tools.Request
			Request.Method = "update"
			Request.ID = tools.StringToInt(Args[1])
			Request.Status = Args[0]
			err := tools.TodoLC(Request)
			if err != nil {
				fmt.Println(err)
				counter = 0
			}
		} else {
			fmt.Println("Invalid command")

		}

	}
}
