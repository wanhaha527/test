FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test1
ADD . $GOPATH/src/awesomeProject/test1
RUN go build .
EXPOSE 8070
ENTRYPOINT ["./test1"]
