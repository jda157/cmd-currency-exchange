.PHONY: build
build:
	go build -o bin/currency cmd/consoleMoneyRate/main.go

.PHONY: clean
clean:
	rm -rf bin


