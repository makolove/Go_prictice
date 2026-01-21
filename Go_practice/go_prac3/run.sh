#!/bin/bash
BINARY_NAME="my_go_app"

# 增加一步：进入正确的代码目录
cd go_prac3 

echo "开始在 go_prac3 目录下构建..."
go fmt main.go
go build -o $BINARY_NAME main.go

if [ $? -eq 0 ]; then
    echo "构建成功，启动程序..."
    ./$BINARY_NAME
else
    echo "构建失败！"
fi