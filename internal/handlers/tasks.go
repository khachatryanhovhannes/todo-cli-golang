package handlers

import (
	"fmt"
	"sync"
	"todo-cli/internal/models"
	"todo-cli/internal/storage"
)

var (
	tasks  []models.Task
	nextID int
	mu     sync.Mutex
)

func LoadTasks() {
	mu.Lock()
	defer mu.Unlock()

	var err error
	tasks, err = storage.LoadTasks()
	if err == nil && len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}
}

func SaveTasks() {
	err := storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func AddTask(name string) {
	mu.Lock()
	defer mu.Unlock()

	task := models.Task{ID: nextID, Name: name, Done: false}
	tasks = append(tasks, task)
	nextID++
	go SaveTasks()
	fmt.Println("Task added:", name)
}

func ListTasks() {
	mu.Lock()
	defer mu.Unlock()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %s (ID: %d)\n", status, task.Name, task.ID)
	}
}

func DeleteTask(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			go SaveTasks()
			fmt.Println("Task deleted:", id)
			return
		}
	}
	fmt.Println("Task not found.")
}

func MarkTaskAsDone(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			go SaveTasks()
			fmt.Println("Task marked as done:", id)
			return
		}
	}
	fmt.Println("Task not found.")
}

func MarkTaskAsUndone(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = false
			go SaveTasks()
			fmt.Println("Task marked as not done:", id)
			return
		}
	}
	fmt.Println("Task not found.")
}
