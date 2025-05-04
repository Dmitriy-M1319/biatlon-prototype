.PHONY: build clean

build:
	go build -o prototype  cmd/biatlon-prototype/main.go
clean:
	rm prototype
