package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name           string      `json:"name" gorm:"comment:分类名称"`
	Alias          string      `json:"alias" gorm:"comment:分类别名"`
	Description    string      `json:"description" gorm:"type:text;comment:分类描述"`
	Article        []*Article  `json:"article" gorm:"many2many:article_category;comment:此分类文章"`
	ParentCategory []*Category `json:"parent_category" gorm:"many2many:category_relationships;association_jointable_foreignkey:parent_category_id;comment:父目录"`
}
