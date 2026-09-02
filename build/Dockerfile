FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o xapp -ldflags="-s -w" -trimpath . 

FROM alpine:3.22.2

RUN addgroup -g 1000 pax && \
    adduser -u 1000 -G pax -s /bin/sh -D pax

COPY --from=builder /app/xapp /usr/local/bin/xapp

RUN chmod +x /usr/local/bin/xapp

USER pax

WORKDIR /workspace
