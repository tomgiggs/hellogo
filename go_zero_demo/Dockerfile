FROM ccr.deepwisdomai.com/infra/go-base-image/go-builder:1.0-alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

LABEL stage=gobuilder

# go的环境变量
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV GOPRIVATE *.deepwisdomai.com
ENV GO111MODULE on

WORKDIR /build/dw

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/pinggo ./pinggo.go


FROM ccr.deepwisdomai.com/pub/ci/alpine
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/pinggo /app/pinggo
COPY --from=builder /app/etc /app/etc

CMD ["./pinggo"]
