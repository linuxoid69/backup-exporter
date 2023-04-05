# Makefile
.PHONY:	build push

PREFIX = reg.my-itclub.ru/monitoring
IMAGE = backup-exporter
VERSION = 0.1.1
TAG = $(VERSION)

# build_app:
# 	CGO_ENABLED=0 go build -o rootfs/opt/backup-exporter cmd/backup-exporter/main.go

build:
	docker build --platform linux/amd64 --build-arg VERSION=$(VERSION) --pull -t $(PREFIX)/$(IMAGE):$(TAG) .

push:
	docker push $(PREFIX)/$(IMAGE):$(TAG)
