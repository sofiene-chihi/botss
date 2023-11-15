build:
	go build -o build/go-api

run-db-containers:
	cd container && docker-compose up -d