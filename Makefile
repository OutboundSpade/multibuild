all: build-linux build-darwin build-windows
BINARY_NAME=multibuild
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/$(BINARY_NAME)-amd64 -v .
	GOOS=linux GOARCH=386 go build -o bin/linux/$(BINARY_NAME)-386 -v .
	GOOS=linux GOARCH=arm go build -o bin/linux/$(BINARY_NAME)-arm -v .
	GOOS=linux GOARCH=arm64 go build -o bin/linux/$(BINARY_NAME)-arm64 -v .

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/$(BINARY_NAME)-amd64 -v .
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/$(BINARY_NAME)-arm64 -v .

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/$(BINARY_NAME)-amd64.exe -v .
	GOOS=windows GOARCH=386 go build -o bin/windows/$(BINARY_NAME)-386.exe -v .
	GOOS=windows GOARCH=arm go build -o bin/windows/$(BINARY_NAME)-arm.exe -v .
	GOOS=windows GOARCH=arm64 go build -o bin/windows/$(BINARY_NAME)-arm64.exe -v .

clean:
	rm -rf bin
	go clean -modcache