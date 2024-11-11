# gin-django-example

#### 介绍
```
仿django项目布局，gin通过接口实现注入分层
app->(user|xxx)->handler -> service ->repository
   为什么这样做是因为，我是重度django玩儿家，习惯django那种项目模块隔离，不同的模块功能归类。
这样代码结构更加清晰，希望对你有帮助。
```

#### 软件架构
```
go1.22+
mysql5.7
gin、gorm、swagger
```

#### 安装教程
```
go mod tidy
go mod vendor
swag init  # swagger文档
```
#### 启动
```
1.  go build main.go
2.  sudo docker build Dockerfile .
3.  sudo docker-compose up -d
```
#### 目录结构
```
gin-django-example
|-app                   # 模块合集
|--user                 # 模块
|---handler             # 模块下的对外API 
|---model               # 模块下定义的struct
|---repository          # 模块下调用数据库操作
|---service             # 模块下实现业务逻辑
|---router.go           # 模块下统一注册router
|--app.go               # 定义swagger路由、启动端口
|--server.go            # 实现服务优雅启动/关闭
|-cmd
|-docs
|-exception
|-logs
|-middleware            # 中间件、拦截器、token身份认证
|-pkg                   # 第三方、自有包初始化
|--db
|--sentry
|-utils                 # 一些通用的工具
|-main.go               # 主程序入口

```
#### 使用说明

```
1. http://localhost:8000/docs/index.html # 文档地址

```

#### 参与贡献
1.  
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
