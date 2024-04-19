run:
	# GIN_MODE=release go run ./cmd/web/main.go
	go run ./cmd/web/main.go

swag:
	swag init -g ./cmd/web/main.go