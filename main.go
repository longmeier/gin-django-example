package main

import (
	"context"
	"fmt"
	"gin-django-example/app"
	"gin-django-example/pkg/db"
	"gin-django-example/pkg/eye"
	"gin-django-example/pkg/jwt"
	"gin-django-example/pkg/log"
	"gin-django-example/pkg/redis"
	"gin-django-example/pkg/sentry"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func GetPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
func ReadConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

func main() {
	// 读取配置
	configPath := filepath.Join(GetPwd(), "config.yaml")
	env := ReadConfig(configPath)
	// 错误日志收集
	sentry.NewSentry(env)
	// redis
	redis.NewRedis(env)
	// 日志
	logger := log.NewLogger(env)
	// 数据库
	sql := db.NewSql(env)
	gormDB := db.NewDB(env, logger)
	// jwt salt-key
	jwt.NewJwtKey(env)
	// app注册
	appParam := eye.AppParam{
		Db: gormDB, Sql: sql, Env: env, Log: logger,
	}
	// 创建并初始化 App 实例
	app := app.NewApp(&appParam)
	// 启动应用
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		logger.Error(fmt.Sprintf("启动服务器失败: %v", err))
	}

	// 优雅关闭服务器
	defer func() {
		if err := app.Stop(ctx); err != nil {
			logger.Error(fmt.Sprintf("关闭服务器失败: %v", err))
		}
	}()
}
