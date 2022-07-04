# 构建阶段 可指定阶段别名
# 基础镜像
FROM golang:1.18.1-bullseye AS builder

# 容器环境变量设置，会覆盖默认的变量值
ENV GO111MODULE="on"
ENV	GOPROXY="https://goproxy.cn,direct"

# 作者
LABEL author="vjman"
LABEL email="vjman@foxmail.com"

# 工作区，可以是本地文件（绝对路径）或者git仓库
WORKDIR /home/lkserver/project

# 复制源文件（本地或git仓库）到容器里
COPY . .

# 编译可执行二进制文件 
# RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app app.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app app.go

# 构建生产镜像，使用alpine当基础镜像，使用scratch时，vjman多次尝试启动失败
FROM debian:bullseye-slim AS prod

# 工作区
WORKDIR /webserver

# 解决证书问题
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && update-ca-certificates

# 在构建阶段复制时区
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 在构建阶段复制可执行的go二进制文件
COPY --from=builder /home/lkserver/project/app .

# 在构建阶段复制配置文件
COPY --from=builder /home/lkserver/project/config ./config

# 配置端口
EXPOSE 80
 
# 启动服务
ENTRYPOINT ["./app"]

# multi-stage build
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds