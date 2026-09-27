package main

import (
	"fmt"
	"os"
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

}
