# 第一阶段：构建阶段
FROM golang:1.20 AS builder

# 设置 Go 环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /home/workspace

# 复制源代码
ADD . /home/workspace

# 下载依赖并构建应用
RUN go mod tidy
RUN go build -o web main.go

# 第二阶段：运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 复制编译后的二进制文件
COPY --from=builder /home/workspace/web .

# 确保配置文件存在并复制到容器中
COPY ./initialize/config /root/initialize/config
COPY ./logs /root/logs

# 声明服务端口
EXPOSE 8080
EXPOSE 6066

# 运行应用
ENTRYPOINT ["./web"]
