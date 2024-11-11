# 基础镜像
FROM golang:alpine

# 作者
MAINTAINER longmeier

# 工作目录 执行go命令的目录
WORKDIR $GOPATH/src/paycenter

# 将本地内容添加到镜像指定目录
COPY . $GOPATH/src/paycenter

# 运行
RUN go env -w GO111MODULE=auto

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build grpcbuf/server/server.go

# 指定镜像内部服务监听的端口
EXPOSE 9101

# 镜像默认入口命令
ENTRYPOINT ["./server"]