FROM golang:alpine as builder
#FROM golang:1.16 AS builder

ARG APP_RELATIVE_PATH
COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
	&& mkdir -p bin/ \
    && go build -o ./bin/ ./...

FROM alpine AS runner
ARG APP_RELATIVE_PATH

WORKDIR /src/app/${APP_RELATIVE_PATH}/cmd

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin/ .
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/configs/ ../configs/

# user 4 order 2 bingfood 0 cart 1
EXPOSE 8002
EXPOSE 9002
VOLUME /data/conf

#CMD ["./cmd"]