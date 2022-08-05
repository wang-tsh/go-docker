# Compile stage
FROM golang:1.18 AS build-env
ENV GO111MODULE=on \
    GOPROXY="https://mirrors.aliyun.com/goproxy/"
ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -o /dockerdev/server
# Final stage
FROM debian:buster
USER root
EXPOSE 8080
WORKDIR /
COPY --from=build-env /dockerdev/server /
COPY --from=build-env /dockerdev/config /config
CMD ["/server"]