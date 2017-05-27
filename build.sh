#!/bin/bash

rm -rf bin && mkdir bin && cd bin && cp ../sysgo.reg . && cp ../sysgo.bat .
echo "clean and make bin"

echo "build sysgo (mac x64) ..."
GOOS=darwin GOARCH=amd64 go build ../. && tar -zcvf sysgo_darwin_x64.tar.gz sysgo && rm -rf sysgo

echo "build sysgo (windows x86) ..."
GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui" ../. && zip -r sysgo_windows_x86.zip sysgo.exe sysgo.reg sysgo.bat && rm -rf sysgo.exe

echo "build sysgo (windows x64) ..."
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" ../. && zip -r sysgo_windows_x64.zip sysgo.exe sysgo.reg sysgo.bat && rm -rf sysgo.exe

echo "build sysgo (linux x86) ..."
GOOS=linux GOARCH=386 go build ../. && tar -zcvf sysgo_linux_x86.tar.gz sysgo && rm -rf sysgo

echo "build sysgo (linux x64) ..."
GOOS=linux GOARCH=amd64 go build ../. && tar -zcvf sysgo_linux_x64.tar.gz sysgo && rm -rf sysgo

rm -rf conf sysgo.bat sysgo.reg

echo "build all arch Successfully"
