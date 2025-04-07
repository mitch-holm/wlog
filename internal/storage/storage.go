package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mholm/wlog/internal/task"
)

const (
	storageDir  = ".wlog"
	storageFile = "tasks.json"
)

// Storage handles persistence of tasks
type Storage struct {
	filePath string
}

// NewStorage creates a new storage instance
func NewStorage() (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	storagePath := filepath.Join(homeDir, storageDir)
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return nil, err
	}

	return &Storage{
		filePath: filepath.Join(storagePath, storageFile),
	}, nil
}

// SaveTasks saves the given tasks to disk
func (s *Storage) SaveTasks(tasks []*task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

// LoadTasks loads tasks from disk
func (s *Storage) LoadTasks() ([]*task.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*task.Task{}, nil
		}
		return nil, err
	}

	var tasks []*task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
} 