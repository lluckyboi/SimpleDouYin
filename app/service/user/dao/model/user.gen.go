// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	UserID        int64  `gorm:"column:user_id;primaryKey" json:"user_id"` // id
	Username      string `gorm:"column:username;not null" json:"username"`
	Password      string `gorm:"column:password;not null" json:"password"`
	Name          string `gorm:"column:name;not null" json:"name"`
	FollowCount   int64  `gorm:"column:follow_count;not null" json:"follow_count"`     // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;not null" json:"follower_count"` // 粉丝总数
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
