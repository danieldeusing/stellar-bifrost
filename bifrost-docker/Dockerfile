FROM golang:stretch as builder
ENV BUILD_FROM_SRC 1

COPY /bifrost /bifrost/

ADD install.sh /
RUN chmod +x /install.sh
RUN /install.sh

# -----------------------
FROM debian:latest

COPY --from=builder /built-bifrost /go/bin/bifrost

ADD build-config /usr/bin/build-config
RUN chmod +x /usr/bin/build-config

RUN ["mkdir", "-p", "/opt/bifrost"]

ADD entry.sh /entry.sh
RUN chmod +x /entry.sh

ADD apk-server /apk-server
RUN chmod +x /apk-server
RUN /apk-server

ADD configs /configs

ENTRYPOINT ["/entry.sh"]

EXPOSE 8800
