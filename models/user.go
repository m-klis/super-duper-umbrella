package models

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Umur      int       `json:"umur"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"create_at"`
}

func (i *User) Bind(r *http.Request) error {
	if i.Nama == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
