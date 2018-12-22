#源镜像
FROM golang:latest
#作者
MAINTAINER liuyc "liuyc@knowbox.cn"
#设置工作目录
WORKDIR $GOPATH/src/github.com/Quons/go-gin-example
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/Quons/go-gin-example
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 9082
#最终运行docker的命令
ENTRYPOINT  ["./go-gin-example"]
