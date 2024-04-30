FROM golang:1.22.0 AS builder
WORKDIR /usr/local/bin
COPY . .
RUN make build

FROM alpine:latest
WORKDIR /usr/local/bin
COPY --from=builder /usr/local/bin/url-shortner-application ./app-go
COPY --from=builder /usr/local/bin/env ./env
EXPOSE 8080
CMD ["./app-go"]