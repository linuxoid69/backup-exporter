# Makefile
.PHONY:	build push

PREFIX = linuxoid69/backup-exporter
IMAGE = backup-exporter
VERSION = 0.1.1
TAG = $(VERSION)

build:
	docker build --platform linux/amd64 --build-arg VERSION=$(VERSION) --pull -t $(PREFIX)/$(IMAGE):$(TAG) .

push:
	docker push $(PREFIX)/$(IMAGE):$(TAG)
