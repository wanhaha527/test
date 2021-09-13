FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test12
ADD . $GOPATH/src/awesomeProject/test12
ENV  GOPROXY=https://goproxy.cn
RUN go build .
EXPOSE 8012
ENTRYPOINT ["./test12"]
