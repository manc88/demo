.PHONY: proto test us down up bup mock 

us:
	go run ./cmd/user_service/
lint:
	golangci-lint run
test:
	go test ./... -count=1 -covermode=count -coverpkg ./... -coverprofile=tprof.out 
	go tool cover --html=tprof.out
down:
	docker compose down
up:
	docker compose up
bup:
	docker compose build
	docker compose up
proto:
	protoc  \
    -I=./api/ \
	--go_out=. --go_opt=paths=import  \
    --go-grpc_out=. --go-grpc_opt=paths=import  \
    ./api/*.proto
mock:
	mockgen \
		-source ./internal/user_service/contract.go \
		-destination ./internal/user_service/mocks/mock.go \
		-package mocks
