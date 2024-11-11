package repository

import (
	"database/sql"
	"fmt"
	"gin-django-example/app/user/model"
	"gin-django-example/pkg/eye"
	"gin-django-example/pkg/eye/time"
	"gin-django-example/pkg/log"
	"gorm.io/gorm"
)

// UserRepository 定义用户数据访问接口
type UserRepository interface {
	GetUserByID(id uint) (*model.User, error)
	GetUserInfo(id uint) (*UserInfo, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

// UserRepositoryImpl 是 UserRepository 的实现
type UserRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
	log    *log.Logger
}

type UserInfo struct {
	Id        int       `json:"id"`
	Username  string    `json:"user_name"` // 账户
	Name      string    `json:"name"`      // 姓名
	Mobile    string    `json:"mobile"`    // 手机号
	LastLogin time.Time `json:"last_login"`
	Enabled   int       `json:"enabled"`
	Visual    int       `json:"visual"`
	Created   time.Time `json:"created"`
	UpdateAt  time.Time `json:"update_at"`
	AddUserID int       `json:"add_user_id"`
}

// NewUserRepository 创建 UserRepository 实例
func NewUserRepository(par *eye.AppParam) UserRepository {
	return &UserRepositoryImpl{gormDB: par.Db, sqlDB: par.Sql, log: par.Log}
}

func UserToUserInfo(user *model.User) *UserInfo {
	return &UserInfo{
		Id:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Mobile:    user.Mobile,
		LastLogin: user.LastLogin,
		Enabled:   int(user.Enabled),
		Visual:    int(user.Visual),
		Created:   user.Created,
		UpdateAt:  user.UpdateAt,
		AddUserID: int(user.AddUserID),
	}
}

// GetUserByID 根据用户 ID 获取用户信息
func (r *UserRepositoryImpl) GetUserByID(id uint) (*model.User, error) {
	fmt.Println("repository", 1)
	var user model.User
	if err := r.gormDB.First(&user, id).Error; err != nil {
		fmt.Println("GetUserByID:", err)
		return nil, err
	}
	fmt.Println("user:", &user)
	return &user, nil
}

// GetUserInfo 获取用户信息
func (r *UserRepositoryImpl) GetUserInfo(id uint) (*UserInfo, error) {
	user, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return UserToUserInfo(user), nil
}

// CreateUser 创建新用户
func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return r.gormDB.Create(user).Error
}

// UpdateUser 更新用户信息
func (r *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return r.gormDB.Save(user).Error
}

// DeleteUser 删除用户
func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	return r.gormDB.Delete(&model.User{}, id).Error
}
