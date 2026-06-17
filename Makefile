BINARY_NAME := update-version-tool

build:
	go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)

help:
	@echo "Use 'make build' to build the binary"
	@echo "Use 'make clean' to remove generated files."
