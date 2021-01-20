package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string `json:"title" gorm:"type:text;comment:标题"`
	Content      string `json:"content" gorm:"type:longtext;comment:博文内容"`
	Views        int    `json:"views" gorm:"comment:浏览数"`
	CommentCount int    `json:"comment_count" gorm:"comment:评论数"`
	LikeCount    uint   `json:"like_count" gorm:"comment:点赞数"`
	WriterID     int    `json:"user_id" gorm:"column:writer_id;comment:用户id"`
	Writer       User   `json:"writer"`
	LabelID      uint   `json:"label_id" gorm:"column:label_id;comment:标签"`
	CategoryID   uint   `json:"category_id" gorm:"column:category_id;comment:分类"`
	CommentID    uint   `json:"comment_id" grom:"column:comment_id;comment:评论"`
}
