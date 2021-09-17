package domain

import "time"

type User struct {
	ID        int
	Name      string `validate:"required,max=255"`
	Email     string `validate:"required,email"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (User) IsNode() {}

func (t *User) BeforeSave() error {
	return validator.Struct(t)
}
