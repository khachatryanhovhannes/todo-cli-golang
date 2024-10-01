package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-cli/internal/handlers"
)

func main() {
	handlers.LoadTasks()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command:\nA: Add\nL: List\nDE: Delete\nDO: Done\nUD: Undone\nE: Exit\n")
		fmt.Print("_______________________________________________________\n")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToUpper(cmd))

		switch cmd {
		case "A":
			fmt.Print("Enter task name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			handlers.AddTask(name)
		case "L":
			handlers.ListTasks()
		case "DE":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			handlers.DeleteTask(id)
		case "DO":
			fmt.Print("Enter task ID to mark as done: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			handlers.MarkTaskAsDone(id)
		case "UD":
			fmt.Print("Enter task ID to mark as not done: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			handlers.MarkTaskAsUndone(id)
		case "E":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
