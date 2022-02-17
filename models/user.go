package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Age       int       `json:"age" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFilter struct {
	Name    string
	AgeUp   int
	AgeDown int
	Status  string
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"nama"`
	Age       int    `json:"umur"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// func (i *User) Bind(r *http.Request) error {
// 	if i.Nama == "" {
// 		return fmt.Errorf("name is a required field")
// 	}
// 	return nil
// }

// func (*User) Render(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
