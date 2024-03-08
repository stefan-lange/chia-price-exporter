FROM curlimages/curl as downloader-task
ARG TASK_VERSION=3.35.1
ARG TARGETARCH
WORKDIR /download
RUN curl --retry 6 -Ls https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_${TARGETARCH}.tar.gz | tar -xz -C /download/ \
    && chmod +x /download/task

# -------------------------------------
# builder
# -------------------------------------
FROM golang:1.22-alpine3.18 as builder
RUN apk add --update --no-cache \
        bash \
        build-base
# install task
COPY --from=downloader-task /download/task /usr/local/bin/task
WORKDIR /build
# cache dependencies (for faster builds)
COPY go.* ./
RUN go mod download
# build
COPY . .
RUN task build

FROM alpine:3 as builder-ssl-certs
RUN apk add --no-cache ca-certificates
# -------------------------------------
# prod
# -------------------------------------
FROM scratch as prod
COPY --from=builder-ssl-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/bin/chia-price-exporter /usr/local/bin/chia-price-exporter
ENTRYPOINT ["/usr/local/bin/chia-price-exporter"]
