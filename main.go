package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func getTaskById(tsk []Task, id int) (*Task, bool) {

	for i := range tsk {
		if tsk[i].ID == id {
			return &tsk[i], true
		}
	}
	return nil, false
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
	case "done":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid id %d %v \n", id, err)
			os.Exit(1)
		}
		task, ok := getTaskById(tasks, id)
		if !ok {
			fmt.Fprintf(os.Stderr, "Task not found for id %d = %v \n", id, ok)
			os.Exit(1)
		} else {
			task.Done = true
			tasksByte, err := json.Marshal(tasks)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error encoding task %v\n", err)
				os.Exit(1)
			}

			if err := os.WriteFile("tasks.json", tasksByte, 0644); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to file %v\n", err)
				os.Exit(1)
			}
		}

	default:
		fmt.Println("Unknown command")
	}
}
