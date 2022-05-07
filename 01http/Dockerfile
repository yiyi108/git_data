FROM golang:1.18.1 AS build
WORKDIR /root/gopath/src/zhuoye/01http/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /root/gopath/src/zhuoye/01http /root/gopath/src/zhuoye/01http
EXPOSE 80
ENV ENV local
WORKDIR /root/gopath/src/zhuoye/01http
ENTRYPOINT ["./httpserver"]
