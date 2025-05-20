run:
	docker-compose up -d
	go run ./cmd/profile/main.go

migrate:
	go run cmd/migrator/main.go

goose-add:
	goose -dir migrations create addTestData sql

rebuild:
	docker-compose down
	docker-compose up -d --build
	go run ./cmd/profile/main.go