FROM golang:1.23-alpine
MAINTAINER OceanWang
WORKDIR /app
ENV XDEMO_CONFIG_SOURCE=container
# 单独复制mod和sum文件，去下载依赖
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
# 拷贝当前git目录所有内容到/app下
COPY . .
RUN go build -o demo .
RUN ls -al & pwd
# APP 访问端口
EXPOSE 7077
# Xdemo Server启动执行的命令
CMD ["/app/demo -tlscert=/ocean.crt -tlskey='ocean.key' -tlsport=17443"]