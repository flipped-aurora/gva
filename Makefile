GVA = "gva"
GfVueAdmin = "gf-vue-admin"
GfVueAdminMysqlMac = "gf-vue-admin-mysql-mac"
GfVueAdminMysqlMacM1 = "gf-vue-admin-mysql-mac-m1"
GfVueAdminMysqlLinux = "gf-vue-admin-mysql-linux"
GfVueAdminMysqlWindows = "gf-vue-admin-mysql-windows.exe"
GfVueAdminPostgresMac = "gf-vue-admin-postgres-mac"
GfVueAdminPostgresMacM1= "gf-vue-admin-postgres-mac-m1"
GfVueAdminPostgresLinux = "gf-vue-admin-postgres-linux"
GfVueAdminPostgresWindows = "gf-vue-admin-postgres-windows.exe"

gfva:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.io,direct
	go build -o ${GfVueAdmin} cmd/gfva/main.go
	@if [ -f ${GfVueAdmin} ] ; then ./${GfVueAdmin} initdb && rm ${GfVueAdmin} ; fi

env:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.io,direct
	go get -u github.com/flipped-aurora/gva@master
	@gva initdb -f gf

gf-vue-admin-mysql-mac:
	@if [ -f ${GfVueAdminMysqlMac} ] ; then rm ${GfVueAdminMysqlMac} ; fi
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags "mysql" -o ${GfVueAdminMysqlMac} cmd/main.go

gf-vue-admin-mysql-mac-m1:
	@if [ -f ${GfVueAdminMysqlMacM1} ] ; then rm ${GfVueAdminMysqlMacM1} ; fi
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags "mysql" -o ${GfVueAdminMysqlMacM1} cmd/main.go

gf-vue-admin-mysql-linux:
	@if [ -f ${GfVueAdminMysqlLinux} ] ; then rm ${GfVueAdminMysqlLinux} ; fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags "mysql" -o ${GfVueAdminMysqlLinux} cmd/main.go

gf-vue-admin-mysql-windows:
	@if [ -f ${GfVueAdminMysqlWindows} ] ; then rm ${GfVueAdminMysqlWindows} ; fi
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags "mysql" -o ${GfVueAdminMysqlWindows} cmd/main.go

gf-vue-admin-postgres-mac:
	@if [ -f ${GfVueAdminPostgresMac} ] ; then rm ${GfVueAdminPostgresMac} ; fi
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags "postgres" -o ${GfVueAdminPostgresMac} cmd/main.go

gf-vue-admin-postgres-mac-m1:
	@if [ -f ${GfVueAdminPostgresMacM1} ] ; then rm ${GfVueAdminPostgresMacM1} ; fi
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags "postgres" -o ${GfVueAdminPostgresMacM1} cmd/main.go

gf-vue-admin-postgres-linux:
	@if [ -f ${GfVueAdminPostgresLinux} ] ; then rm ${GfVueAdminPostgresLinux} ; fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags "postgres" -o ${GfVueAdminPostgresLinux} cmd/main.go

gf-vue-admin-postgres-windows:
	@if [ -f ${GfVueAdminPostgresWindows} ] ; then rm ${GfVueAdminPostgresWindows} ; fi
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags "postgres" -o ${GfVueAdminPostgresWindows} cmd/main.go

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go

swagger:
	@gf swagger

check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f ${GfVueAdmin} ] ; then rm ${GfVueAdmin} ; fi
	@if [ -f ${GfVueAdminMysqlMac} ] ; then rm ${GfVueAdminMysqlMac} ; fi
	@if [ -f ${GfVueAdminMysqlLinux} ] ; then rm ${GfVueAdminMysqlLinux} ; fi
	@if [ -f ${GfVueAdminMysqlWindows} ] ; then rm ${GfVueAdminMysqlWindows} ; fi

help:
	@echo "make - 构建gfva终端工具并初始化数据,初始化数据后删除gfva终端工具,启动server项目"
	@echo "make gfva - 构建gfva终端工具 并初始化数据"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 main.go"
	@echo "make clean - 移除二进制文件"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"