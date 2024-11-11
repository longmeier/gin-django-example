package app

import (
	"context"
	"gin-django-example/app/user"
	_ "gin-django-example/docs" // 引入 Swagger 生成的 docs 包
	"gin-django-example/middleware"
	"gin-django-example/pkg/eye"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App 代表整个应用程序的结构体
type App struct {
	httpServer *Server
}

// NewApp 创建并初始化整个应用
func NewApp(par *eye.AppParam) *App {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	s := NewServer(
		gin.Default(),
		par.Log,
		WithServerHost("0.0.0.0"),
		WithServerPort(8000),
	)
	s.Use(middleware.LogMiddleware())

	s.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := s.Group("/api/v1") // middleware.JWTAuth()
	user.NewUserApp(par, v1)
	// 返回初始化完成的 App
	return &App{httpServer: s}
}

// Start 启动 HTTP 服务器
func (app *App) Start(ctx context.Context) error {
	return app.httpServer.Start(ctx)
}

// Stop 停止 HTTP 服务器
func (app *App) Stop(ctx context.Context) error {
	// 调用 Server 的 Stop 方法来停止服务器
	return app.httpServer.Stop(ctx)
}
