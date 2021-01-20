package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    int      `json:"user_id" gorm:"column:user_id;comment:用户id"`
	User      User     `json:"user"`
	Article   Article  `json:"article" gorm:"comment:文章"`
	LikeCount uint     `json:"like_count" gorm:"comment:点赞数"`
	Content   string   `json:"content" gorm:"type:text;comment:评论内容"`
	ReplyID   *uint    `json:"reply_id" gorm:"column:reply_id;comment:回复id"`
	Reply     *Comment `json:"reply" gorm:"comment:回复"`
}
