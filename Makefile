# Internal variables
AIR_PORT := 2345
APP_PORT := 8080

SRC_DIR := ./src/
AIR_DIR := ../.air.toml
TEMP_DIR := ./tmp

BUILD_DIR := ./build/
BUILD_FILE_NAME := engine

# Targets
.PHONY: help
help:
	@printf "%s\n" "Useful targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  make %-15s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Start app by invoking targets, "terminate-running" and "run-air"
	@$(MAKE) terminate-running run-air

.PHONY: run-air
run-air: ## Start Air by using config file
	cd $(SRC_DIR) && air -c $(AIR_DIR)

.PHONY: terminate-running
terminate-running: ## Terminate the running process
	lsof -i:$(AIR_PORT),$(APP_PORT) | grep LISTEN | awk -F " " ' { print $$2 } ' | xargs kill -9

.PHONY: build
build: ## Build source code to executable file
	go build -o $(BUILD_DIR)$(BUILD_FILE_NAME) $(SRC_DIR)

.PHONY: clean
clean: ## Recursively clean tmp and build folder
	rm -rf $(TEMP_DIR) $(BUILD_DIR)