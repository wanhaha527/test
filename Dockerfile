FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test3
ADD . $GOPATH/src/awesomeProject/test3
RUN go build .
EXPOSE 8003
ENTRYPOINT ["./test3"]
