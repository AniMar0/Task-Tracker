package main

import (
	"Task-Tracker/tools"
	"fmt"
	"os"
	"strings"
)

func main() {
	tools.MakeToDoFile()
	counter := 0
	for {
		Args := os.Args[1:]
		if counter == 0 {
			fmt.Println("Please provide a command or type \"help\" for more information")
			counter++
		}
		if Args[0] == "exit" {
			break
		} else if Args[0] == "help" {
			fmt.Print("Commands: \n",
				"list - Lists all tasks \n",
				"list [todo|in-progress|done] - Lists all tasks of a specific status \n",
				"mark-[in-progress|done|todo] [ID] - Marks a task as [in-progress|done|todo] \n",
				"add [Description] - Adds a task \n",
				"update [ID] [Description]- Updates a task \n",
				"delete [ID] - Deletes a task \n",
				"exit - Exits the program \n")
		} else if Args[0] == "list" {
			var Request tools.Request
			if len(Args) == 2 {
				Request.Method = Args[0]
				if len(Args) == 2 {
					Request.List = Args[1]
				} else {
					Request.List = "list"
				}
				err := tools.TodoLC(Request)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else if Args[0] == "add" {
			var Request tools.Request
			if len(Args) == 2 {
				Request.Method = Args[0]
				Request.Description = Args[1]
				err := tools.TodoLC(Request)
				if err != nil {
					fmt.Println(err)
					return
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
					return
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
					return
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
				return
			}
		} else {
			fmt.Println("Invalid command")

		}

	}
}
