# Task-Tracker

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Description

A simple command-line task tracker application written in Go. It allows users to add, update, list, and delete tasks, with each task having a unique ID, description, status (todo, in-progress, done), and timestamps for creation and last update. The tasks are stored in a JSON file for persistence.

## Features

- Automatically creates `db.json` for storage.
- List all tasks or filter by status.
- Add new tasks with trimmed descriptions.
- Update existing tasks by ID.
- Delete tasks by ID.
- Mark tasks as `todo`, `in-progress`, or `done`.
- Responsive table output in the terminal with dynamic column widths.
- Stores timestamps for creation and last update.

## Requirements

- Go 1.23+ recommended
- Windows, macOS, or Linux

## Installation

Clone the repository and download modules:

```powershell
git clone https://github.com/AniMar0/Task-Tracker.git
cd Task-Tracker
go mod tidy
```

## Running the application

Start the CLI:

```powershell
go run main.go
```

You will see the prompt:

Please provide a command or type "help" for more information

>

## Commands

- `list`  
  Lists all tasks.

- `list [status]`  
  Lists tasks filtered by status (`todo`, `in-progress`, `done`).

- `add "Description"`  
  Adds a new task with the given description. Extra spaces are trimmed.
  Example:

  > add "Learn Go basics"

- `update [ID] "Description"`  
  Update a task's description by ID. Example:

  > update 1 "Learn Go CLI basics"

- `delete [ID]`  
  Delete a task by ID. Example:

  > delete 1

- `mark-todo [ID]` / `mark-in-progress [ID]` / `mark-done [ID]`  
  Set task status by ID. Example:

  > mark-done 1

- `help`  
  Show available commands.

- `exit`  
  Exit the program.

## Example session

```
> add "Learn Go basics"
> list
> mark-done 1
> update 1 "Learn Go CLI basics"
> delete 1
> exit
```

## Data storage

All tasks are stored in `db.json` in the project directory. The file is created automatically on first run if it doesn't exist.

## Notes

- Descriptions are trimmed of extra whitespace.
- Table columns adapt to content for clean display.
- IDs are integers assigned incrementally.

## Contributing

Contributions, issues, and feature requests are welcome. Please open an issue or submit a pull request.

## License

This project is open-source and free to use.
