// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFollow = "follow"

// Follow mapped from table <follow>
type Follow struct {
	UID       int64 `gorm:"column:uid;not null" json:"uid"`               // 关注者id
	TargetUID int64 `gorm:"column:target_uid;not null" json:"target_uid"` // 被关注者id
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
