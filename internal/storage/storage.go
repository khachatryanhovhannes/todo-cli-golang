package storage

import (
	"encoding/json"
	"os"
	"todo-cli/internal/models"
)

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}
