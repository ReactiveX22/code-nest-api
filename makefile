BINARY_NAME=code-nest-api.exe
BUILD_DIR=bin

SOURCE=main.go

build:
	@echo "Building the application..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE)

run: build
	@echo "Running the application..."
	$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)