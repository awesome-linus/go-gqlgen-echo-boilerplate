package domain

import "time"

type Task struct {
	ID        int
	Title     string `validate:"required,max=255"`
	UserID    int    `validate:"required"`
	Notes     string `validate:"omitempty,max=65535"`
	Completed bool
	Due       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (Task) IsNode() {}

func (t *Task) BeforeSave() error {
	return validator.Struct(t)
}
