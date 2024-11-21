.PHONY: proto

discovery-serve:
	@go run cmd/discovery/main.go

discovery-build:
	@go build -o bin/discovery cmd/discovery/main.go

bidding-serve:
	@go run cmd/bidding/main.go

bidding-build:
	@go build -o bin/bidding cmd/bidding/main.go

docker-build:
	@go build -o bin/server cmd/$(SERVICE)/main.go
