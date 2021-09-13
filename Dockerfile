FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test16
ADD . $GOPATH/src/awesomeProject/test16
ENV  GOPROXY=https://goproxy.cn
RUN go build .
EXPOSE 8016
ENTRYPOINT ["./test16"]
