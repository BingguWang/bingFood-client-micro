FROM golang:alpine as builder
#FROM golang:1.16 AS builder

COPY . /src
WORKDIR /src/app/user/service/cmd/server

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && mkdir -p bin/  \
    && go build  -o ./server ./...

FROM alpine:latest


WORKDIR /src/app/user/service/cmd/server

#COPY --from=builder /src/wait-for-it.sh /bin/
COPY --from=builder /src/app/user/service/cmd/server/server .
COPY --from=builder /src/app/user/service/configs/ ../../configs/

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

#CMD ["./server", "-conf", "/data/conf"]
#CMD ["./server"]