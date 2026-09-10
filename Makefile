# Переменные для сборки
PROTO_ROOT=api/proto
PROTO_PATH=api/proto/tiny-profile-service/v1
PROTO_OUT=pkg/api/grpc/
OPEN_API_OUT=pkg/api/http/auth/v1
MODULE_NAME=github.com/ElfAstAhe/tiny-profile-service
SERVER_BINARY_NAME=tiny-profile-service
SERVER_BUILD_DIR=./cmd/tiny-profile-service
VERSION=1.0.0
BUILD_TIME=$(shell date +'%Y/%m/%d_%H:%M:%S')
STAGE=DEV

.PHONY: help gen-proto gen-swagger gen-http-client gen-mocks build run test bench static-check clean update-deps

help:
	@echo "Доступные команды для сборки и тестирования:"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'
	@echo "------------------------------------------------------------------------"

# Генерация gRPC кода
gen-proto: ## Сгенерировать gRPC код (Go & gRPC) из Protobuf файлов
	mkdir -p $(PROTO_OUT)
	protoc \
        -I $(PROTO_ROOT) \
		--proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		--go_opt=default_api_level=API_OPAQUE \
		$(PROTO_PATH)/*.proto

# Генерация swagger
gen-swagger: ## Сгенерировать Swagger-документацию (swag init)
	swag init \
		-g $(SERVER_BUILD_DIR)/main.go \
		--parseDependency \
		--parseInternal \
		--exclude ./pkg/api \
		-o docs \
		--parseDepth 3
#	swag init -g cmd/server/main.go

# Генерация http client
gen-http-client: ## Сгенерировать HTTP-клиент на основе swagger.json
	mkdir -p $(OPEN_API_OUT)
	swagger generate client -f ./docs/swagger.json -A tiny-profile-service -t $(OPEN_API_OUT)

# Генерирует моки для интерфейсов в указанной папке, см. {project_root}/.mockery.yml конфиг
gen-mocks: ## Сгенерировать моки для интерфейсов (mockery)
	mockery

# Сборка проекта с прокидыванием переменных
build: gen-proto gen-swagger gen-http-client gen-mocks ## Полная сборка: генерация всего кода + компиляция бинарника
	go build -ldflags "-X '$(MODULE_NAME)/internal/config.AppVersion=$(VERSION)' \
	-X '$(MODULE_NAME)/internal/config.AppBuildTime=$(BUILD_TIME)'" \
	-o ./bin/$(SERVER_BINARY_NAME) $(SERVER_BUILD_DIR)/main.go

#	go build -ldflags "-X '$(MODULE_NAME)/internal/app/client/config.Version=$(VERSION)' \
#    -X '$(MODULE_NAME)/internal/app/client/config.Stage=$(STAGE)' \
#	-X '$(MODULE_NAME)/internal/app/client/config.BuildTime=$(BUILD_TIME)'" \
#	-o ./bin/$(CLIENT_BINARY_NAME) $(CLIENT_BUILD_DIR)/main.go

# Запуск проекта (сначала соберет, потом запустит)
run: build ## Собрать проект и запустить бинарник с локальными флагами (БД, логи)
	./bin/$(SERVER_BINARY_NAME) \
        --log-level "debug" \
		--http-address "localhost:8082" \
		--grpc-address "localhost:50052" \
		--db-driver "postgres" \
		--db-dsn "postgres://svc_profile:password@localhost:5432/test?sslmode=disable&search_path=auth_db" \
		--auth-jwt-secret "jwt-key" \
		--app-cipher-key "12345" \
		--app-token-issuer "tiny-auth-service" \
		--app-max-list-limit 500 \

# Запуск тестов
test: gen-proto gen-mocks ## Запустить модульные и интеграционные тесты проекта
	go test -v ./...

# Запуск бенчмарков (сюда добавляем все вызовы) или разные параметры под один пакет
bench: gen-proto gen-mocks ## Запустить утилиты с замером памяти
#	go test -bench=BenchmarkManager_FullCycle -benchmem ./pkg/infra/cache/test/...
	go test -bench=. -benchmem ./...

# Запуск static check
static-check: ## Запустить статический анализ кода (пропуская автогенерируемый pkg/api)
	staticcheck $$(go list ./... | grep -vE "pkg/api|cmd/grpc-client-test")

# Запуск линтера
lint: ## Запустить линтер revive (пропуская автогенерируемый код)
	revive $$(go list ./... | grep -vE "pkg/api|cmd/gen-tz")

# Очистка бинарников
clean: ## Очистить скомпилированные файлы из папки ./bin
	rm -rf ./bin/*

# обновление зависимостей
update-deps: ## Принудительно обновить и скачать все Go-зависимости проекта
	go get -u -x all

#