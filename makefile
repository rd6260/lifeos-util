BINARY=lifeos-util
PREFIX=$(HOME)/.local/bin

build:
	@go build -o $(BINARY) .

run: build
	@./$(BINARY) $(ARGS)

dev:
	@go run . $(ARGS)

install: build
	@mkdir -p $(PREFIX)
	@cp $(BINARY) $(PREFIX)/$(BINARY)
	@echo "installed to $(PREFIX)/$(BINARY)"

uninstall:
	@rm -f $(PREFIX)/$(BINARY)
	@echo "removed $(PREFIX)/$(BINARY)"

clean:
	@rm -f $(BINARY)
