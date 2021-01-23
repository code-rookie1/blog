package response

import "blog/model"

type Category struct {
	Category model.Category `json:"category"`
}
