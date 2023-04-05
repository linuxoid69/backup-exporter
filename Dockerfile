FROM golang:1.18-alpine3.15 AS builder

WORKDIR /go/src/app

COPY ./ ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o app cmd/backup-exporter/main.go

FROM reg.my-itclub.ru/docker/base/alpine:latest

COPY --from=builder /go/src/app/app /opt/backup-exporter

USER 0

RUN apk update && \
    apk upgrade && \
    chmod a+x /opt/backup-exporter

USER 100500

EXPOSE 9925

CMD ["/opt/backup-exporter"]
