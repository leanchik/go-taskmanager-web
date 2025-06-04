package storage

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

type Storage interface {
	LoadTasks() ([]Task, error)
	SaveTasks([]Task) error
}

type JSONStorage struct {
	FileName string
}

func (s JSONStorage) LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(s.FileName)
	if err != nil {

		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (s JSONStorage) SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, data, 0644)
}
