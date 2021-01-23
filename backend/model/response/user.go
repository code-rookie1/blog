package response

import "blog/model"

type User struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
