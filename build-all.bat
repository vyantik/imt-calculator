@echo off
echo Building for Windows...
go build -o app.exe

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o app-linux

echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o app-mac

echo Done!