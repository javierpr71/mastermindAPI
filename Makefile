
BINARY_NAME=mastermind

build:
	go build -o bin/${BINARY_NAME} main.go

build-all:
	echo "Building for every OS and Platform"	
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-win.exe main.go

clean:
	go clean
	rm bin/*

run:
	go run main.go

all: build-all

docker:
	docker build -t mastermind:latest .
