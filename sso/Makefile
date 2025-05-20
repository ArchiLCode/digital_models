run:
	docker-compose up -d
	go run ./cmd/sso/main.go

migrate:
	go run cmd/migrator/main.go

goose-add:
	goose -dir migrations create add_refresh_table sql

rebuild:
	docker-compose down
	docker-compose up --build