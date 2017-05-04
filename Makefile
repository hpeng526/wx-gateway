export PATH := $(GOPATH)/bin:$(PATH)
export GO15VENDOREXPERIMENT := 1

all: build

build: gateway backend

gateway:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/gateway/wx-gateway_darwin_amd64 ./cmd/gateway/
	env GOOS=linux GOARCH=386 go build -o ./build/gateway/wx-gateway_linux_386 ./cmd/gateway/
	env GOOS=linux GOARCH=amd64 go build -o ./build/gateway/wx-gateway_linux_amd64 ./cmd/gateway/
	env GOOS=linux GOARCH=arm go build -o ./build/gateway/wx-gateway_linux_arm ./cmd/gateway/
	env GOOS=windows GOARCH=386 go build -o ./build/gateway/wx-gateway_windows_386.exe ./cmd/gateway/
	env GOOS=windows GOARCH=amd64 go build -o ./build/gateway/wx-gateway_windows_amd64.exe ./cmd/gateway/
	env GOOS=linux GOARCH=mips64 go build -o ./build/gateway/wx-gateway_linux_mips64 ./cmd/gateway/
	env GOOS=linux GOARCH=mips64le go build -o ./build/gateway/wx-gateway_linux_mips64le ./cmd/gateway/

backend:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/backend/wx-backend_darwin_amd64 ./cmd/backend/
	env GOOS=linux GOARCH=386 go build -o ./build/backend/wx-backend_linux_386 ./cmd/backend/
	env GOOS=linux GOARCH=amd64 go build -o ./build/backend/wx-backend_linux_amd64 ./cmd/backend/
	env GOOS=linux GOARCH=arm go build -o ./build/backend/wx-backend_linux_arm ./cmd/backend/
	env GOOS=windows GOARCH=386 go build -o ./build/backend/wx-backend_windows_386.exe ./cmd/backend/
	env GOOS=windows GOARCH=amd64 go build -o ./build/backend/wx-backend_windows_amd64.exe ./cmd/backend/
	env GOOS=linux GOARCH=mips64 go build -o ./build/backend/wx-backend_linux_mips64 ./cmd/backend/
	env GOOS=linux GOARCH=mips64le go build -o ./build/backend/wx-backend_linux_mips64le ./cmd/backend/

clean:
	rm -rf ./build/
