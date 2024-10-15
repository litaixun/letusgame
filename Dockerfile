# 使用官方 Go 镜像
FROM golang:1.19

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译 Go 程序
RUN go build -o letusgame .

# 暴露端口
EXPOSE 12345

# 启动应用
CMD ["./letusgame"]

