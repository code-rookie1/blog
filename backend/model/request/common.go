package request

type PageInfo struct {
	Page     int `json:"page" validate:"required"`
	PageSize int `json:"pageSize" validate:"required"`
}

type GetByID struct {
	Id int `json:"id" validate:"required"`
}
