proto:
	./scripts/gen-proto.sh
	
swag:
	swag init -g api/api.go --parseDependency

run: swag proto
	go run cmd/main.go

.PHONY: swag, proto, run
