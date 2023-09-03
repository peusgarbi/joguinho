DEFAULT_GOAL := default

APPLICATION_NAME ?= joguinho
BUILD_PATH ?= build
ENTRY_FILE ?= main.go

.PHONY: build # Builda os executáveis
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_PATH)/$(APPLICATION_NAME) $(ENTRY_FILE) && \
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_PATH)/$(APPLICATION_NAME).exe $(ENTRY_FILE)
