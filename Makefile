PWD = $(shell pwd)
# Название приложения
APP_NAME = cube_cli
PATH = "$(PATH):$(PWD)/bin"

# Запуск сервиса
.PHONY: bin
bin:
	mkdir -p $(PWD)/bin
	go build -o $(PWD)/bin/$(APP_NAME)  $(PWD)/cmd/$(APP_NAME)/main.go
	export PATH

# Запустить тесты
.PHONY: test
test:
	go test $(PWD)/... -v

