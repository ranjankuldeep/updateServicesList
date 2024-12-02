APP_NAME = update_serverList_data
BUILD_DIR = ./build
MAIN_FILE = ./main.go

update: build run

build:
	@echo "Building the application..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

run:
	@echo "Running the application..."
	$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up build files..."
	rm -rf $(BUILD_DIR)

help:
	@echo "Available commands:"
	@echo "  make update   - Build and run the application"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the compiled binary"
	@echo "  make clean    - Remove built files"
