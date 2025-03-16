# 使用多阶段构建以优化镜像大小
FROM golang:1.21-alpine as builder

# 设置构建时的工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级的alpine作为运行时基础镜像
FROM alpine:latest

# 安装ca-certificates用于HTTPS支持
RUN apk add --no-cache ca-certificates tzdata

# 设置时区为上海
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录
WORKDIR /app

# 从builder阶段复制编译好的应用和配置文件
COPY --from=builder /app/main /app/
COPY --from=builder /app/settings.yaml /app/

# 暴露应用端口
EXPOSE 8080

# 启动应用
CMD ["/app/main"]