PWD = $(shell pwd)
# Название приложения
APP_NAME = cube_cli

.DEFAULT_GOAL := bin

PATH  := $(PATH):$(PWD)/bin
SHELL := env PATH=$(PATH) /bin/bash

# Запуск сервиса
.PHONY: bin
bin:
	mkdir -p $(PWD)/bin
	go build -o $(PWD)/bin/$(APP_NAME)  $(PWD)/cmd/$(APP_NAME)/main.go

# Запустить тесты
.PHONY: test
test:
	go test $(PWD)/... -v

