APP=BaXingNengLiang

all: build
build:
	go build -o build/${APP}-mac cmd/*.go
linux:
	export GOARCH=amd64;export GOOS=linux;go build -o build/${APP}-linux cmd/*.go
arm:
	export GOARCH=arm;export GOOS=linux;go build -o build/${APP}-arm cmd/*.go
win:
	export GOARCH=amd64;export GOOS=windows;go build -o build/${APP}-win.exe cmd/*.go
run:
	./build/${APP}
clean:
	rm -rf ./build/${APP}
.PHONY: build

