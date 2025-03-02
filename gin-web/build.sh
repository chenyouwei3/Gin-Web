#!/bin/bash

# 设置变量
IMAGE_NAME="my-gin-web"
VERSION="v1.0.0"
DOCKERFILE="Dockerfile"

# 打印日志
echo "🚀 开始构建 Docker 镜像: $IMAGE_NAME:$VERSION"

# 运行构建命令
docker build -t $IMAGE_NAME:$VERSION -f $DOCKERFILE .

# 检查是否构建成功
if [ $? -eq 0 ]; then
    echo "✅ 镜像构建成功: $IMAGE_NAME:$VERSION"
else
    echo "❌ 镜像构建失败，请检查错误日志！"
    exit 1
fi

# 运行镜像（可选）
echo "🔥 运行容器："
docker run -d -p 8080:8080 -p 6066:6066 --name my-gin-container $IMAGE_NAME:$VERSION

echo "✅ 容器已启动，访问: http://localhost:8080"
