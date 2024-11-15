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

migrate-up:
	@echo "Running Migrations..."
	go build -o bin/migrate.exe pkg/migrate/migrate.go
	bin/migrate.exe --up

migrate-down:
	@echo "Running Migrations..."
	go build -o bin/migrate.exe pkg/migrate/migrate.go
	bin/migrate.exe --down
