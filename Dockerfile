FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test15
ADD . $GOPATH/src/awesomeProject/test15
ENV  GOPROXY=https://goproxy.cn
RUN go build .
EXPOSE 8015
ENTRYPOINT ["./test15"]
