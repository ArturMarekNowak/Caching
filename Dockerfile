FROM golang:1.22

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -v -o ./caching ./cmd/app
EXPOSE 8080
CMD ["./caching"]