FROM golang:latest
MAINTAINER me
WORKDIR $GOPATH/src/awesomeProject/test2
ADD . $GOPATH/src/awesomeProject/test2
RUN go build .
EXPOSE 8070
ENTRYPOINT ["./test2"]
