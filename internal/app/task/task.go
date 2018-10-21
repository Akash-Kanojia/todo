package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Task represents an activity or a note.
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserEmail string    `json:"user_email"`
	Starred   bool      `json:"starred"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTask return instance of the task entity.
func NewTask(title, body, email string) (task Task, err error) {
	if len(body) < 0 {
		err = fmt.Errorf("task : body cannot be empty")
		return
	}

	task.ID = uuid.New().String()
	task.Title = title
	task.Body = body
	task.UserEmail = email
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return
}
