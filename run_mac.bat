@REM 构建mac可执行文件
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags -H=windowsgui