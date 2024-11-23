# 第一阶段：构建阶段
FROM golang:1.20-alpine AS builder

# 设置 Go 环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有代码复制到容器中
COPY . .

# 下载依赖并编译应用
RUN go mod tidy && go build -o /dist/app .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装需要的依赖（如 CA certificates，很多应用需要它们）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到当前镜像
COPY --from=builder /dist/app .

# 声明服务端口
EXPOSE 8888

# 启动容器时执行应用
CMD ["/root/app"]
