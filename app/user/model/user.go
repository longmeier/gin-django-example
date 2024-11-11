package model

import (
	"gin-django-example/pkg/eye/time"
)

const TableNameUser = "user"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string    `gorm:"column:username;comment:账户" json:"username"` // 账户
	Name      string    `gorm:"column:name;comment:姓名" json:"name"`         // 姓名
	Mobile    string    `gorm:"column:mobile;comment:手机号" json:"mobile"`    // 手机号
	Password  string    `gorm:"column:password" json:"password"`
	Salt      string    `gorm:"column:salt" json:"salt"`
	LastLogin time.Time `gorm:"column:last_login;default:null" json:"last_login"`
	Enabled   int       `gorm:"column:enabled;default:1" json:"enabled"`
	Visual    int       `gorm:"column:visual;default:1" json:"visual"`
	Created   time.Time `gorm:"column:created;not null;default:CURRENT_TIMESTAMP" json:"created"`
	UpdateAt  time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP" json:"update_at"`
	AddUserID int       `gorm:"column:add_user_id" json:"add_user_id"`
}

func (*User) TableName() string {
	return TableNameUser
}
