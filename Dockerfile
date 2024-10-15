# 使用官方 Go 语言镜像
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将当前目录的内容复制到工作目录
COPY . .

# 编译 Go 应用
RUN go build -o letusgame .

# 使用更小的镜像来运行程序
FROM alpine:latest

# 安装 Fyne 依赖
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从 builder 镜像复制编译好的二进制文件
COPY --from=builder /app/letusgame .

# 运行应用
CMD ["./letusgame"]
