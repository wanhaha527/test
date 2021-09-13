FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test5
ADD . $GOPATH/src/awesomeProject/test5
RUN go build .
EXPOSE 8005
ENTRYPOINT ["./test5"]
