# syntax=docker/dockerfile:1

## Build
FROM golang:1.21-bookworm AS build
LABEL org.opencontainers.image.source https://github.com/hadret/dht22_exporter

WORKDIR /go/src
COPY . /go/src
RUN CGO_ENABLED=0 go build -a -o dht22_exporter


## Deploy
FROM gcr.io/distroless/base-debian12

COPY --from=build /go/src/dht22_exporter /dht22_exporter

USER nonroot:nonroot

EXPOSE 10005
ENTRYPOINT [ "/dht22_exporter" ]
