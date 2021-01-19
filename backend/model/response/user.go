package response

import "blog/model"

type User struct {
	User model.User `json:"user"`
}
