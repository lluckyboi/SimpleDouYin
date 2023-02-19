#获取golang
FROM golang:1.19.0-alpine3.16 as go

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /server

# 将代码复制到容器中
COPY etc /server/etc
COPY internal /server/internal
COPY ./chat.go /server

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
CMD ["chmod 733 chat.go","chmod +x chat.go"]
# 启动容器时运行的命令
CMD ["go run chat.go -f etc/chat.aml"]