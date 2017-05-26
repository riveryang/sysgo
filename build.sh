#!/bin/bash

rm -rf bin && mkdir bin && cd bin && cp -R ../conf . && cp ../sysgo.reg . && cp ../sysgo.bat .
echo "clean and make bin"

echo "build sysgo (mac x64) ..."
GOOS=darwin GOARCH=amd64 go build ../. && tar -zcvf sysgo_darwin_amd64.tar.gz sysgo conf && rm -rf sysgo

echo "build sysgo (windows x86) ..."
GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui" ../. && tar -zcvf sysgo_windows_386.tar.gz sysgo.exe conf sysgo.reg sysgo.bat && rm -rf sysgo.exe

echo "build sysgo (windows x64) ..."
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" ../. && tar -zcvf sysgo_windows_amd64.tar.gz sysgo.exe conf sysgo.reg sysgo.bat && rm -rf sysgo.exe

echo "build sysgo (linux x86) ..."
GOOS=linux GOARCH=386 go build ../. && tar -zcvf sysgo_linux_386.tar.gz sysgo conf && rm -rf sysgo

echo "build sysgo (linux x64) ..."
GOOS=linux GOARCH=amd64 go build ../. && tar -zcvf sysgo_linux_amd64.tar.gz sysgo conf && rm -rf sysgo

rm -rf conf sysgo.bat sysgo.reg

echo "build all arch Successfully"
