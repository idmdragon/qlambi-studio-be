package user

import "time"

type UserFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatUser(user User, token string) UserFormatter {
	return UserFormatter{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
