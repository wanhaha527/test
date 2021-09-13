FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test17
ADD . $GOPATH/src/awesomeProject/test17
ENV  GOPROXY=https://goproxy.cn
RUN go build .
EXPOSE 8017
ENTRYPOINT ["./test17"]
