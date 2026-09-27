package main

import (
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: todo <add|list|done> [task]")
		os.Exit(1)
	}
	fmt.Println(args[1])
	taskOne := Task{
		ID:    1,
		Title: "Title one",
	}
	tasks := []Task{
		taskOne,
	}

	tasks = append(tasks, Task{
		ID:    2,
		Title: "Title two",
	})

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo <add> <title>")
			os.Exit(1)
		}
		title := strings.Join(args[2:], " ")
		fmt.Println("adding a task")
		fmt.Println(title)
		tasks = append(tasks, Task{
			ID:    len(tasks) + 1,
			Title: title,
		})
		fmt.Println(tasks)
	case "list":
		fmt.Println("listing tasks")
	default:
		fmt.Println("Unknown command")
	}
}
