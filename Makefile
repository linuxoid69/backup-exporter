# Makefile
.PHONY:	build push

PREFIX = linuxoid69
IMAGE = backup-exporter
TAG = 0.1.1

build:
	docker build --platform linux/amd64 --pull -t $(PREFIX)/$(IMAGE):$(TAG) .

push:
	docker push $(PREFIX)/$(IMAGE):$(TAG)
