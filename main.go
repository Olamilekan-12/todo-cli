package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: todo <add|list|done> [task]")
		os.Exit(1)
	}
	var tasks []Task
	tasksByte, err := os.ReadFile("tasks.json")
	if errors.Is(err, os.ErrNotExist) {
		// No file to work with atm...
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %v\n", err)
		os.Exit(1)
	} else {
		err = json.Unmarshal(tasksByte, &tasks)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshalling file %v\n", err)
			os.Exit(1)
		}
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo <add> <title>")
			os.Exit(1)
		}
		title := strings.Join(args[2:], " ")
		fmt.Println("adding a task")

		tasks = append(tasks, Task{
			ID:    len(tasks) + 1,
			Title: title,
		})
		data, err := json.Marshal(tasks)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding task %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile("tasks.json", data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file %v\n", err)
			os.Exit(1)
		}
	case "list":
		fmt.Println("listing tasks")
		for _, task := range tasks {
			box := "[ ]"
			if task.Done {
				box = "[x]"
			}
			fmt.Printf("%d. %s %s\n", task.ID, box, task.Title)
		}
	default:
		fmt.Println("Unknown command")
	}
}
