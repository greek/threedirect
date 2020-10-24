default: run

build:
	go build -o server.exe server.go

run:
	$(MAKE) build
	server.exe