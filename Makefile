MAIN: ./cmd/main.go
	@ go run ./cmd/main.go

build: ./cmd/main.go | builds
	go build -o builds/out cmd/main.go
	./builds/out


builds:
	mkdir -p $@

