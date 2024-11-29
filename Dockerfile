FROM golang:1.22 as builder

# 配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

#工作目录
WORKDIR /app
ADD .  /app

#在Docker工作目录下执行命令
RUN go build -o main

FROM ubuntu:22.04 as runner

#作者信息
MAINTAINER "mols"

ENV FOO_ENV=production

WORKDIR /app

COPY --from=builder /app/main ./
COPY --from=builder /app/envs/ ./envs

#暴露端口
EXPOSE 18005

#执行项目的命令
CMD ["/app/main"]


