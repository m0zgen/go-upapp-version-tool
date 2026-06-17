BINARY_NAME := update-version-tool

build:
	go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME)

help:
	@echo "Используйте 'make build' для сборки утилиты."
	@echo "Используйте 'make clean' для удаления сгенерированных файлов."
