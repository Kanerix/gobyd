.PHONY: proto

serve:
	@go run cmd/bidding/main.go

build:
	@go build -o bin/bidding cmd/bidding/main.go

docker-build:
	@go build -o bin/server cmd/$(SERVICE)/main.go
