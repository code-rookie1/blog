package response

import "blog/model"

type Comment struct {
	Comment model.Comment `json:"comment"`
}
