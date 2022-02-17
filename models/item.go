package models

import (
	"time"
)

type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	Price       int       `json:"price" validate:"required"`
}

type ItemFilter struct {
	StartDate *time.Time
	EndDate   *time.Time
	Name      string
	Page      int
	View      int
}

type ItemResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Price       int    `json:"price"`
}

// type ItemRequest struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	CreatedAt   string `json:"created_at"`
// }

// func (i *Item) Bind(r *http.Request) error {
// 	if i.Name == "" {
// 		return fmt.Errorf("name is a required field")
// 	}
// 	return nil
// }

// func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
