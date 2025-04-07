package task

import (
	"time"
)

// Importance represents the importance level of a task (1-5)
type Importance int

const (
	MinImportance Importance = 1
	MaxImportance Importance = 5
)

// Task represents a single work item
type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Importance  Importance `json:"importance"`
	CreatedAt   time.Time `json:"created_at"`
	Tags        []string  `json:"tags,omitempty"`
}

// NewTask creates a new task with the given description and importance
func NewTask(description string, importance Importance) *Task {
	// Ensure importance is within valid range
	if importance < MinImportance {
		importance = MinImportance
	}
	if importance > MaxImportance {
		importance = MaxImportance
	}

	return &Task{
		ID:          generateID(),
		Description: description,
		Importance:  importance,
		CreatedAt:   time.Now(),
		Tags:        make([]string, 0),
	}
}

// generateID creates a unique ID for the task
func generateID() string {
	return time.Now().Format("20060102150405")
} 