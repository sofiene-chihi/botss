build:
	go build -o build/go-api

run-redis:
	docker run --name redis-ecommerce-instance -p 4000:6379 -d redis