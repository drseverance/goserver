# ---------- Build Stage ----------
FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goserver .

# ---------- Runtime Stage ----------
FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y wget && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/goserver .

ENV PORT=8991

EXPOSE 8991

CMD ["./goserver"]
