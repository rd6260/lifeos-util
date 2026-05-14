BINARY=lifeos-util

build:
	@go build -o $(BINARY) .

run: build
	@./$(BINARY)

clean:
	@rm -f $(BINARY)

dev:
	@go run .
