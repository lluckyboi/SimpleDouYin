// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	UserID     int64     `gorm:"column:user_id;not null" json:"user_id"`
	CommentID  int64     `gorm:"column:comment_id;primaryKey" json:"comment_id"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	CreateDate time.Time `gorm:"column:create_date;not null" json:"create_date"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
