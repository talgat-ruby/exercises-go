FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o myapp .

FROM alpine:latest

RUN apk add --no-cache libpq

COPY --from=builder /app/myapp /app/myapp

ENTRYPOINT ["/app/myapp"]

EXPOSE 8080