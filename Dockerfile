
FROM golang:1.8.3-alpine3.6
MAINTAINER Bin Jiang <jiangbin@haier.com>

# repo
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN echo "http://mirrors.aliyun.com/alpine/v3.6/main/" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.6/community/" >> /etc/apk/repositories

# timezone
RUN apk update
RUN apk add --no-cache tzdata \
    && echo "Asia/Shanghai" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# Add Tini
# RUN apk add --no-cache tini
# ENTRYPOINT ["/sbin/tini", "--"]

# move to GOPATH
RUN mkdir -p /go/src/git.haier.net/monitor/promext-apm-plugin
COPY . $GOPATH/src/git.haier.net/monitor/promext-apm-plugin
WORKDIR $GOPATH/src/git.haier.net/monitor/promext-apm-plugin


# build
RUN go build -o /app/run main.go

WORKDIR /app
CMD ["/app/run"]