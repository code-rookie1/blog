package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name          string     `json:"name" gorm:"comment:分类名称"`
	Alias         string     `json:"alias" gorm:"comment:分类别名"`
	Description   string     `json:"description" gorm:"type:text;comment:分类描述"`
	Article       []Article  `json:"article" gorm:"comment:此分类文章"`
	SubCategoryID *uint      `json:"sub_category_id" gorm:"comment:子分类id"`
	SubCategory   []Category `json:"sub_category" gorm:"foreignkey:SubCategoryID" gorm:"子分类"`
}
