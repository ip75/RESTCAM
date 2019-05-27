rem build for linux/amd64
rem linux/amd64

set GOOS=linux
set GOARCH=amd64

rem for linux/386
rem set GOARCH=386

go build
