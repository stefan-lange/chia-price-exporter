FROM golang:1.20-alpine as builder-golang-base
RUN apk add --update --no-cache \
        bash \
        build-base
WORKDIR /build

FROM builder-golang-base as builder
# cache dependencies (for faster builds)
COPY go.* ./
RUN go mod download
# build
COPY . .
RUN make build

FROM scratch
COPY --from=builder /build/bin/chia-price-exporter /usr/local/bin/chia-price-exporter
ENTRYPOINT ["/usr/local/bin/chia-price-exporter"]
