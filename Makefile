.PHONY: all build run gotool clean help

GVA = "gva"

all: init gva

init:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.io,direct

gva: init
	go build -o ${GVA} cmd/gva/main.go
	@if [ -f ${GVA} ] ; then mv ${GVA} $GOPATH/bin && rm ${GVA} ; fi

linux-build: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${GVA} cmd/gva/main.go

windows-build: init
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${GVA}.exe cmd/gva/main.go

mac-build: init
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${GVA} cmd/gva/main.go

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f ${GVA} ] ; then rm ${GVA} ; fi

help:
	@echo "make - 构建gva终端工具并初始化数据,初始化数据后删除gva终端工具"
	@echo "make gva - 构建gva终端工具并移动gva到GOPATH/bin目录下"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make clean - 移除二进制文件"