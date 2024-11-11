package app

import (
	"context"
	"fmt"
	"gin-django-example/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
}
type Option func(s *Server)

func NewServer(engine *gin.Engine, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}

	// 初始化 http.Server 实例
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s, // 使用 Gin 引擎作为 HTTP 处理器
	}

	return s
}
func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}
func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// Start 启动 HTTP 服务器
func (s *Server) Start(ctx context.Context) error {
	// 启动 HTTP 服务器
	fmt.Println("服务启动->:0.0.0.0:", s.port)
	if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server failed to start:", err)
		return err
	}
	return nil
}

// Stop 停止 HTTP 服务器
func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Shutting down server...")

	// 创建一个超时上下文（最大超时 5 秒）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务器
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
		return err
	}

	fmt.Println("Server exited gracefully")
	return nil
}
