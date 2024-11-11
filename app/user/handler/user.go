package handler

import (
	"encoding/json"
	"gin-django-example/app/user/model"
	"gin-django-example/app/user/service"
	"gin-django-example/pkg/eye"
	"gin-django-example/pkg/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserHandler 处理与用户相关的 HTTP 请求
type UserHandler struct {
	userService service.UserService
	log         *log.Logger
}

// NewUserHandler 创建 UserHandler 实例
func NewUserHandler(userService service.UserService, log *log.Logger) *UserHandler {
	return &UserHandler{userService: userService, log: log}
}

// @Summary GetUser
// @Description 获取用户信息
// @Tags user
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} User "OK"
// @Failure 400 {object} model.User "Invalid request"
// @Router /api/v1/user/info [get]
func (h *UserHandler) GetUser(ctx *gin.Context) {
	// 获取用户ID
	userId := ctx.Query("id")
	uId, err := strconv.Atoi(userId)
	if err != nil {
		eye.HandleError(ctx, eye.BadRequestCode, "参数错误", nil)
		return
	}
	// service层获取用户信息
	iuser, err := h.userService.GetUser(uint(uId))
	if err != nil {
		eye.HandleSuccess(ctx, eye.SuccessCode, "获取用户信息成功", nil)
		return
	}
	eye.HandleSuccess(ctx, eye.SuccessCode, "OK", iuser)
}

// RegisterUser 注册新用户

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var user model.User
	// 获取操作人
	value, exists := ctx.Get("requestUserId")
	if exists {
		user.AddUserID = value.(int)
	}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		eye.HandleError(ctx, eye.BadRequestCode, "参数格式错误", nil)
		return
	}
	if err := h.userService.RegisterUser(&user); err != nil {
		eye.HandleError(ctx, eye.BadRequestCode, err.Error(), nil)
		return
	}
	eye.HandleSuccess(ctx, eye.SuccessCode, "注册成功", nil)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	//var user model.User
	//if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	//	http.Error(w, "请求数据无效", http.StatusBadRequest)
	//	return
	//}
	//if err := h.userService.UpdateUser(&user); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//json.NewEncoder(w).Encode(user)
}

// RemoveUser 删除用户
func (h *UserHandler) RemoveUser(ctx *gin.Context) {

}
