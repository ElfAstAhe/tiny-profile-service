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

.PHONY: build run test clean

# Генерация gRPC кода
gen-proto:
	mkdir -p $(PROTO_OUT)
	protoc \
        -I $(PROTO_ROOT) \
		--proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		--go_opt=default_api_level=API_OPAQUE \
		$(PROTO_PATH)/*.proto

# Генерация swagger
gen-swagger:
	swag init \
		-g $(SERVER_BUILD_DIR)/main.go \
		--parseDependency \
		--parseInternal \
		--exclude ./pkg/api \
		-o docs \
		--parseDepth 3
#	swag init -g cmd/server/main.go

gen-http-client:
	mkdir -p $(OPEN_API_OUT)
	swagger generate client -f ./docs/swagger.json -A tiny-profile-service -t $(OPEN_API_OUT)

gen-mocks:
# Генерирует моки для всех интерфейсов в указанной папке
	mockery

# Сборка проекта с прокидыванием переменных
#build: gen-proto gen-swagger gen-http-client gen-mocks
build:
#build: gen-proto
	go build -ldflags "-X '$(MODULE_NAME)/internal/config.AppVersion=$(VERSION)' \
	-X '$(MODULE_NAME)/internal/config.AppBuildTime=$(BUILD_TIME)'" \
	-o ./bin/$(SERVER_BINARY_NAME) $(SERVER_BUILD_DIR)/main.go

#	go build -ldflags "-X '$(MODULE_NAME)/internal/app/client/config.Version=$(VERSION)' \
#    -X '$(MODULE_NAME)/internal/app/client/config.Stage=$(STAGE)' \
#	-X '$(MODULE_NAME)/internal/app/client/config.BuildTime=$(BUILD_TIME)'" \
#	-o ./bin/$(CLIENT_BINARY_NAME) $(CLIENT_BUILD_DIR)/main.go

# Запуск проекта (сначала соберет, потом запустит)
run: build
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
test:
	go test -v ./...

# Запуск static check
static-check:
	staticcheck $$(go list ./... | grep -vE "pkg/api|cmd/grpc-client-test")

# Очистка бинарников
clean:
	rm -rf ./bin/*

# обновление зависимостей
update-deps:
	go get -u -x all

#