package models

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (miUser *User) AddUser(id int, name string, createAt time.Time, status bool) {
	miUser.Id = id
	miUser.Name = name
	miUser.CreatedAt = createAt
	miUser.Status = status
}
