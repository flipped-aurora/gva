.PHONY: all build run gotool clean help

BINARY="gva"

all: gotool mac-build

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
