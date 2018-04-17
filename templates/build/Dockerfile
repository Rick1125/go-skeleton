#build stage
FROM docker.senal.club/rhizen/go-builder:latest AS builder

# install xz
RUN apt-get update && apt-get install -y xz-utils \
    && rm -rf /var/lib/apt/lists/*

# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

ARG program
COPY ${program} /go/src/${program}
WORKDIR /go/src/${program}
RUN ln -sf /go/vendor
RUN make install

# strip and compress the binary
RUN strip --strip-unneeded /go/bin/${program}
RUN upx /go/bin/${program}

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
ARG program
ARG port
COPY --from=builder /go/bin/${program} .
COPY config.${program}.yml ./config.yml

CMD ["./${program}", "--config=/app/config.yml", "run"]
EXPOSE ${port}

