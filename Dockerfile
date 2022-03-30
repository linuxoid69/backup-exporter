FROM reg.netcitylife.ru/docker/base/alpine:3.15.0

COPY rootfs /

USER 0

RUN apk update && \
    apk upgrade && \
    chmod a+x /opt/backup-exporter

USER 100500

EXPOSE 9925

CMD ["/opt/backup-exporter"]
