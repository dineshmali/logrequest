hello:
	echo "Hello"

build:
	go build -o bin/main src/main.go

run:
	go run src/main.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm64 go build -o bin/logrequest-linux src/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/logrequest-win src/main.go
all: hello build
