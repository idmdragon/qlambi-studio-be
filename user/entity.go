package user

import "time"

type User struct {
	Id             int
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RegisterUserInput struct {
	Name     string
	Email    string
	Password string
}
