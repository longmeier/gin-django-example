package service

import (
	"errors"
	"fmt"
	"gin-django-example/app/user/model"
	"gin-django-example/app/user/repository"
	"gin-django-example/pkg/encrypt"
	"gin-django-example/pkg/log"
)

// UserService 定义用户服务接口
type UserService interface {
	GetUser(id uint) (*repository.UserInfo, error)
	GetUserDetail(id uint) (*model.User, error)
	RegisterUser(user *model.User) error
	UpdateUser(user *model.User) error
	RemoveUser(id uint) error
}

// UserServiceImpl 是 UserService 的实现
type UserServiceImpl struct {
	userRepository repository.UserRepository
	log            *log.Logger
}

// NewUserService 创建 UserService 实例
func NewUserService(userRepo repository.UserRepository, logger *log.Logger) UserService {
	return &UserServiceImpl{userRepository: userRepo, log: logger}
}

// GetUser 获取用户信息
func (s *UserServiceImpl) GetUser(id uint) (*repository.UserInfo, error) {
	fmt.Println("server", id)
	return s.userRepository.GetUserInfo(id)
}

func (s *UserServiceImpl) GetUserDetail(id uint) (*model.User, error) {
	fmt.Println("server", id)
	return s.userRepository.GetUserByID(id)
}

// RegisterUser 注册新用户
func (s *UserServiceImpl) RegisterUser(user *model.User) error {
	if user.Username == "" {
		return errors.New("账号不能为空")
	}
	var newUser model.User
	newUser.Name = user.Name
	newUser.Username = user.Username
	newUser.Mobile = user.Mobile
	newUser.AddUserID = user.AddUserID
	salt, err := encrypt.GenerateSalt(10)
	if err != nil {
		fmt.Println("生成salt失败:", err)
		return err
	}
	newUser.Salt = salt
	passward, err := encrypt.GeneratePBKDF2Hash(user.Password, salt, 10000)
	if err != nil {
		fmt.Println("生成密码失败:", err)
		return err
	}
	fmt.Println("passward:", passward)
	newUser.Password = passward
	return s.userRepository.CreateUser(&newUser)
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	return s.userRepository.UpdateUser(user)
}

// RemoveUser 删除用户
func (s *UserServiceImpl) RemoveUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}
