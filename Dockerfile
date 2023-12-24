FROM golang:1.21-alpine3.18 AS builder

WORKDIR /go/src/app

COPY ./ ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o app cmd/backup-exporter/main.go

FROM alpine:3.18

COPY --from=builder /go/src/app/app /opt/backup-exporter

RUN adduser -h /opt -S -u 100500 exporter \
    && chmod a+x /opt/backup-exporter

USER 100500

EXPOSE 9925

CMD ["/opt/backup-exporter"]
