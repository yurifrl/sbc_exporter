FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o sbc_exporter

FROM alpine:latest

WORKDIR /
COPY --from=builder /app/sbc_exporter /sbc_exporter

EXPOSE 9110

ENTRYPOINT ["/sbc_exporter"] 