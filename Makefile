BIN_NAME=auto-attendance
LD_FLAGS="-s -w"

all: cp-env build-linux

build-all: build-linux build-windows build-mac #build-web
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/${BIN_NAME} -ldflags=${LD_FLAGS} .
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/${BIN_NAME}_windows.exe -ldflags=${LD_FLAGS} .
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/${BIN_NAME}_mac -ldflags=${LD_FLAGS} .
# build-web:
# 	GOOS=js GOARCH=wasm go build -o bin/${BIN_NAME}_web.wasm -ldflags=${LD_FLAGS} .

cp-env:
	cp .env bin/.

clean:
	go clean
	rm -rf ./bin/*
