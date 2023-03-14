#/bin/sh

if [ "$1" ]
then 
	# 生成 pb 程序
	protoc -I=proto --go_out=./app/ $1.proto
	
	# 注入tag
	protoc-go-inject-tag -input=app/pb/$1.pb.go
else
	echo "Wrong parameter. It should be like this, ./proto2go.sh file_name"
fi