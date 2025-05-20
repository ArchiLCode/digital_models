swag:
	cd ./gateway && swag init -g ./cmd/main.go
rebuild:
	docker-compose down -v
	docker-compose up --build
run:
	docker-compose up
stop:
	docker-compose down -v

# Определяем список директорий для линтинга
LINT_DIRS = gateway profile sso

# Стадия lint
lint:
	@for dir in $(LINT_DIRS); do \
		echo "Linting $$dir..."; \
		cd $$dir && golangci-lint run; \
		cd ..; \
	done