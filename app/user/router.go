package user

import (
	"gin-django-example/app/user/handler"
	"gin-django-example/app/user/repository"
	"gin-django-example/app/user/service"
	"gin-django-example/middleware"
	"gin-django-example/pkg/eye"
	"github.com/gin-gonic/gin"
)

// NewUserApp 创建并初始化整个应用
func NewUserApp(par *eye.AppParam, router *gin.RouterGroup) {
	// 初始化 UserRepository
	userRepo := repository.NewUserRepository(par)

	// 初始化 UserService
	userService := service.NewUserService(userRepo, par.Log)

	// 初始化 UserHandler
	userHandler := handler.NewUserHandler(userService, par.Log)

	UserRouter(router, userHandler)
}

func UserRouter(r *gin.RouterGroup, handler *handler.UserHandler) {

	authRouter := r.Group("/user", middleware.JWTAuth())
	{
		//authRouter.GET("/info", handler.GetUser)
		authRouter.POST("/register", handler.RegisterUser)
	}
	r.GET("/user/info", handler.GetUser)
}
