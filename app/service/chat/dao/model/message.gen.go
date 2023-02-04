// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Content    string    `gorm:"column:content;not null" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	UID        int64     `gorm:"column:uid;not null" json:"uid"`
	TargetUID  int64     `gorm:"column:target_uid;not null" json:"target_uid"`
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
