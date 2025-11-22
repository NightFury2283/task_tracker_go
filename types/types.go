package types

import (
	"time"
)

//файл хранит все структуры и типы для проекта

type StorageData struct {
	Tasks  map[int]Task `json:"tasks"`
	NextID int          `json:"nextId"`
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Status string

const (
	Todo        Status = "todo"
	In_progress Status = "in-progress"
	Done        Status = "done"
)
