FROM golang:1.25-alpine AS builder

ARG SERVICE_NAME

RUN apk update && apk add --no-cache ca-certificates tzdata git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o binary ./cmd/${SERVICE_NAME}

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/binary /app

USER 1000:1000

ENTRYPOINT ["/app"]