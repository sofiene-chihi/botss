build:
	go build -o build/go-api

build-arm64:
	GOOS=linux GOARCH=arm64 go build -o build/go-bot-arm64

run:
	go run main.go

run-db-containers:
	cd container && docker-compose up -d