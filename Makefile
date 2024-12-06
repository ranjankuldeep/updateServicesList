APP_NAME = update_serverList_data
BUILD_DIR = ./build
MAIN_FILE = ./main.go

# Target architecture
GOOS = linux
GOARCH = amd64

update:
	@echo "Building the application for $(GOOS)/$(GOARCH)..."
	mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@if [ -f $(BUILD_DIR)/$(APP_NAME) ]; then \
		echo "Build successful! Running the application..."; \
		$(BUILD_DIR)/$(APP_NAME); \
	else \
		echo "Build failed! Binary not found."; \
		exit 1; \
	fi

build:
	@echo "Building the application for $(GOOS)/$(GOARCH)..."
	mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@if [ -f $(BUILD_DIR)/$(APP_NAME) ]; then \
		echo "Build successful! Binary located at $(BUILD_DIR)/$(APP_NAME)"; \
	else \
		echo "Build failed! Binary not found."; \
		exit 1; \
	fi

run: build
	@echo "Running the application..."
	$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up build files..."
	rm -rf $(BUILD_DIR)

help:
	@echo "Available commands:"
	@echo "  make update   - Build for Linux x86_64 and run the application"
	@echo "  make build    - Build the application for Linux x86_64"
	@echo "  make run      - Build and run the application"
	@echo "  make clean    - Remove built files"
