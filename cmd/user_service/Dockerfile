#build & debug
FROM golang:alpine
WORKDIR /user_service
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/user_service/userservice ./cmd/user_service/main.go
EXPOSE 50051
CMD ["./cmd/user_service/userservice"]