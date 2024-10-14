# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 复制当前目录下的所有文件到容器中的 /app 目录
COPY . .

# 编译 Go 程序
RUN go mod init letusgame && go mod tidy && go build -o letusgame .

# 运行生成的二进制文件
CMD ["./letusgame"]

# 暴露端口
EXPOSE 12345
