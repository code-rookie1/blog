package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID           int         `json:"user_id" gorm:"column:user_id;comment:用户id"`
	User             User        `json:"user"`
	ArticleID        int         `json:"article_id" gorm:"column:article_id;comment:文章id"`
	Article          Article     `json:"article"`
	LikeCount        uint        `json:"like_count" gorm:"comment:点赞数"`
	Content          string      `json:"content" gorm:"type:text;comment:评论内容"`
	ParentComment   []*Comment   `json:"parent_comment" gorm:"many2many:comment_relationships;association_jointable_foreignkey:parent_comment_id;comment:父评论"`
}