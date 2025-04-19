BINARY_NAME=kinotime

MAIN_FILE=cmd/main.go

build:
	go mod tidy
	# gofumpt -l -w .
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run: build
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	go mod tidy
	gofumpt -l -w .
	rm data -rf
	rm -f $(BINARY_NAME)

.PHONY: build run clean help