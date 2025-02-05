# Build stage
FROM golang:1.23-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app
COPY go.mod  ./
RUN go mod download
COPY . .
RUN apk update && apk add --no-cache make
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make all

# Staging 
FROM alpine:latest 
COPY --from=builder /app/receipt-service /usr/local/bin/
RUN mkdir /config
COPY --from=builder /app/config/config.yml /config/config.yml
RUN chmod +x /usr/local/bin/receipt-service
EXPOSE  5555
CMD ["receipt-service"]
