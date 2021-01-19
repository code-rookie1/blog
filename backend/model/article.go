package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string      `json:"title" gorm:"type:text;comment:标题"`
	Content      string      `json:"content" gorm:"type:longtext;comment:博文内容"`
	Views        int         `json:"views" gorm:"comment:浏览数"`
	CommentCount int         `json:"comment_count" gorm:"comment:评论数"`
	LikeCount    uint        `json:"like_count" gorm:"comment:点赞数"`
	UserID       int         `json:"user_id" gorm:"column:user_id;comment:用户id"`
	User         User        `json:"user"`
	Label        []*Label    `json:"label" gorm:"many2many:article_label;comment:标签"`
	Category     []*Category `json:"category" gorm:"many2many:article_category;comment:分类"`
}
