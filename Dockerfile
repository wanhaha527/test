FROM golang:latest
MAINTAINER wangaiqin
WORKDIR $GOPATH/src/testProject
ADD . $GOPATH/src/testProject
#COPY <src> <dest>
ENV GOPROXY=https://goproxy.cn
#/cmd/test
RUN go build ./cmd/test.go
EXPOSE 8080
ENTRYPOINT ["./test"]
