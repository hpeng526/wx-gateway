export PATH := $(GOPATH)/bin:$(PATH)
export GO15VENDOREXPERIMENT := 1

all: build

build: app

app:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/wx-gateway_darwin_amd64 ./web
	env GOOS=linux GOARCH=386 go build -o ./build/wx-gateway_linux_386 ./web
	env GOOS=linux GOARCH=amd64 go build -o ./build/wx-gateway_linux_amd64 ./web
	env GOOS=linux GOARCH=arm go build -o ./build/x-gateway_linux_arm ./web
	env GOOS=windows GOARCH=386 go build -o ./build/wx-gateway_windows_386.exe ./web
	env GOOS=windows GOARCH=amd64 go build -o ./build/wx-gateway_windows_amd64.exe ./web
	env GOOS=linux GOARCH=mips64 go build -o ./build/wx-gateway_linux_mips64 ./web
	env GOOS=linux GOARCH=mips64le go build -o ./build/wx-gateway_linux_mips64le ./web

clean:
	rm -rf ./build/