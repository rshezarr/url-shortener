run-db:
	docker run --name=urls -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres:13-alpine

run-api:
	go run ./cmd/main.go