FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o blog-crud ./cmd

EXPOSE 8080

CMD ["./blog-crud"]
