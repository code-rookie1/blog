package response

import "blog/model"

type Article struct {
	Article model.Article `json:"article"`
}
