package presenter

import "api-test/entity"

type User struct {
	ID        entity.ID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	IsActive  bool      `json:"isActive"`
	Roles     []string  `json:"roles"`
}