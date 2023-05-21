FROM alpine:3.17.3
ARG TARGETARCH
COPY "dist/hello_linux_${TARGETARCH}" /usr/bin/hello
WORKDIR /app
ENTRYPOINT ["hello"]
