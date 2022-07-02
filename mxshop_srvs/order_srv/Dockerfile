FROM golang:1.17 as builder

ARG dir=app

# 构建后以旧生效的变量，CGO_ENABLED禁用cgo 然后指定OS等，启用go module
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /$dir

COPY . .

#并go build
RUN go build -o main .

# 运行时使用alpine作为基础镜像
FROM alpine

# 构建时使用的变量，结束后失效
ARG dir=app

WORKDIR /$dir
# 为了防止代码中请求https链接报错，我们需要将证书纳入到scratch中
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert

COPY --from=builder /$dir .

# 指定运行时环境变量
ENV GIN_MODE=release

ENTRYPOINT ["./main"]
