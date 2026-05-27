package models

import "time"

type User struct {
	ID		int       `json:"ID"`
	Name	string    `json:"name"`
	Email	string    `json:"email"`
	CreatedAt	time.Time `json:"createdAt"`
	UpdateAt	time.Time `json:"updatedAt"`
}