
## Run the application
run:
	@go run cmd/api/main.go



build:
	@echo "Building..."
	
	
	@go build -o tmp/main cmd/api/main.go

load-test:
	@echo "Running load test..."
	@go run tests/load.go

