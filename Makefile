# Makefile для миграций и запуска cmd/main.go

# Получаем DNS базы данных из .config.yml
DATABASE_DSN=$(shell grep 'database_dsn:' .config.yml | awk '{print $$2}')
GOOSE_DIR=migrations
GOOSE_CMD=goose -dir $(GOOSE_DIR) postgres $(DATABASE_DSN)
APP_CMD=go run cmd/main.go
WIRE_DIR=./app/di
SWAGGER_MAIN = ./cmd/main.go

# Проверка состояния миграций
status:
	$(GOOSE_CMD) status

# Применить все миграции
up:
	$(GOOSE_CMD) up

# Откатить последнюю миграцию
down:
	$(GOOSE_CMD) down

# Создать новую миграцию (пример: make create name=название_миграции)
create:
	$(GOOSE_CMD) create $(name) sql

# Применить миграции до определённой версии (пример: make up_to version=номер_версии)
up_to:
	$(GOOSE_CMD) up-to $(version)

# Откатить миграции до определённой версии (пример: make down_to version=номер_версии)
down_to:
	$(GOOSE_CMD) down-to $(version)

# Сборка проекта (создает исполняемый файл в ./bin/app)
build:
	go build -ldflags="-s -w" -o /bin/app cmd/main.go

# Запуск cmd/main.go
run:
	$(APP_CMD)

# Пересборка и запуск cmd/main.go
rebuild_run: build
	./bin/app
# Пересборка и запуск cmd/main.go
wire:
	wire $(WIRE_DIR)
# Пересборка и запуск cmd/main.go
swagger:
	swag init -g $(SWAGGER_MAIN)

