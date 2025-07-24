FROM alpine:latest

WORKDIR /app

RUN apk update && apk add go

COPY ./ /app

ENTRYPOINT sh ./build_go_app.sh
