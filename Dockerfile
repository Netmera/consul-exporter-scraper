FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build .

FROM alpine:3.19.1  

RUN apk --no-cache add ca-certificates curl

RUN curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl

WORKDIR /root/

RUN apk add --no-cache libc6-compat
COPY --from=builder /app/consul-exporter-scraper .

CMD ["./consul-exporter-scraper"]
